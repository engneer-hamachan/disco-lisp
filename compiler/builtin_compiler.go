package compiler

import (
	"disco/base"
	"fmt"
)

type BuiltinCompilerIF interface {
	builtinCompile(
		codes []any,
		s *base.S,
		caller string,
		file_name *string,
		row *int,
	) ([]any, error)

	typePropagation(
		s *base.S,
		caller string,
		file_name *string,
		row *int,
	) error

	getMinArgumentCount() int
	getMaxArgumentCount() int

	getName() string
	setAlias(name string)
}

var BuiltinCompilers = make(map[any]BuiltinCompilerIF)

type BuiltinCompiler struct {
	name               string
	returnType         int
	firstArgumentType  int
	secondArgumentType int
	thirdArgumentType  int
	minArgumentCount   int
	maxArgumentCount   int
}

func (b BuiltinCompiler) getMinArgumentCount() int {
	return b.minArgumentCount
}

func (b BuiltinCompiler) getMaxArgumentCount() int {
	return b.maxArgumentCount
}

func (b BuiltinCompiler) getName() string {
	return b.name
}

func (b *BuiltinCompiler) setAlias(name string) {
	b.name = name
}

func (b BuiltinCompiler) decideFname() string {
	switch b.name {
	case "peek-lambda":
		return "lambda"
	default:
		return b.name
	}
}

func (b BuiltinCompiler) checkArgumentCount(
	s *base.S,
	row *int,
) error {

	fname := b.decideFname()

	if sLength(s) < b.minArgumentCount {
		*row = base.InformationWhenParsing[fname].Row

		return fmt.Errorf(
			"too few arguments %s want %d arguments",
			b.name,
			b.minArgumentCount,
		)
	}

	if sLength(s) > b.maxArgumentCount {
		*row = base.InformationWhenParsing[fname].Row

		return fmt.Errorf(
			"too many arguments %s want %d arguments",
			b.name,
			b.maxArgumentCount,
		)
	}

	return nil
}

func (b BuiltinCompiler) checkArgumentCountIsEven(
	s *base.S,
	row *int,
) error {

	if sLength(s)%2 != 0 {
		*row = base.InformationWhenParsing[b.name].Row

		return fmt.Errorf(
			"do not match arguments %s want even arguments",
			b.name,
		)
	}

	return nil
}

func (b BuiltinCompiler) setFunctionArgumentTypes(
	tstack *base.S,
	caller string,
	argument_type int,
	row *int,
) error {

	current_argument_type, ok :=
		FunctionArgumentTypes[MapKeys{caller, tstack.Val.(string)}]
	if !ok {
		return nil
	}

	switch current_argument_type {
	case base.ANY, base.SYMBOL:
		if tstack.Val.(string)[0] == '&' {
			break
		}

		FunctionArgumentTypes[MapKeys{caller, tstack.Val.(string)}] = argument_type

	case base.OPTIONAL, argument_type:
		break

	default:
		*row = base.InformationWhenParsing[tstack.Val.(string)].Row

		return fmt.Errorf(
			"%s is %v not %v",
			tstack.Val.(string),
			base.TypeToString(current_argument_type),
			base.TypeToString(argument_type),
		)
	}

	return nil
}

func (b BuiltinCompiler) setFunctionReturnTypes(caller string) {
	FunctionReturnTypes[caller] = b.returnType
	FunctionReturnTypes[b.getName()] = b.returnType
}

func (b BuiltinCompiler) setFunctionChoiceReturnTypes(
	caller string,
	choice_type int,
) {

	FunctionReturnTypes[caller] = choice_type
	FunctionReturnTypes[b.getName()] = choice_type
}

func (b BuiltinCompiler) appendFunctionOptionalTypes(
	caller string,
	types int,
) {

	FunctionOptionalTypes[caller] = append(FunctionOptionalTypes[caller], types)

	FunctionOptionalTypes[b.getName()] =
		append(FunctionOptionalTypes[b.getName()], types)
}

func (b BuiltinCompiler) clearFunctionOptionalTypes(caller string) {
	delete(FunctionOptionalTypes, caller)
	delete(FunctionOptionalTypes, b.getName())
}

func (b BuiltinCompiler) isSymbolOrWantType(
	tstack *base.S,
	want_type int,
	row *int,
	caller string,
) (bool, error) {

	want_types_message := base.TypeToString(want_type)

	if want_type == base.ANY {
		return false, nil
	}

	switch tstack.Type {
	case want_type, base.ANY:
		return false, nil

	case base.NIL:
		if want_type == base.LIST {
			return false, nil
		}

		*row = base.InformationWhenParsing[tstack.Val].Row

		return false, b.invalidArgumentErrorMsg(
			tstack.Val,
			base.TypeToString(tstack.Type),
			want_types_message,
		)

	case base.OPTIONAL:
		optional_types := FunctionOptionalTypes[tstack.Val.(string)]

		if base.IsMatchOptionalType(optional_types, want_type) {
			return false, nil
		}

		return false, b.invalidArgumentErrorMsg(
			tstack.Val,
			base.OptionalTypeToString(optional_types),
			want_types_message,
		)

	case base.SYMBOL:
		optional_types, ok :=
			FunctionArgumentOptionalTypes[MapKeys{caller, tstack.Val.(string)}]
		if ok {
			if base.IsMatchOptionalType(optional_types, want_type) {
				return false, nil
			}

			return false, b.invalidArgumentErrorMsg(
				tstack.Val,
				base.OptionalTypeToString(optional_types),
				want_types_message,
			)
		}

		if tstack.IsQuoted == false {
			return true, nil
		}

		*row = base.InformationWhenParsing[tstack.Val].Row

		return false, b.invalidArgumentErrorMsg(
			tstack.Val,
			"quoted symbol",
			want_types_message,
		)

	default:
		*row = base.InformationWhenParsing[tstack.Val].Row

		return false, b.invalidArgumentErrorMsg(
			tstack.Val,
			base.TypeToString(tstack.Type),
			want_types_message,
		)
	}
}

func (b BuiltinCompiler) isSymbolOrNumberType(
	tstack *base.S,
	row *int,
	caller string,
) (bool, error) {

	switch tstack.Type {
	case base.INT, base.FLOAT, base.ANY:
		return false, nil

	case base.OPTIONAL:
		optional_types := FunctionOptionalTypes[tstack.Val.(string)]

		if base.IsMatchOptionalTypeForNumber(optional_types) {
			return false, nil
		}

		return false, b.invalidArgumentErrorMsg(
			tstack.Val,
			base.OptionalTypeToString(optional_types),
			"int or float.",
		)

	case base.SYMBOL:
		optional_types, ok :=
			FunctionArgumentOptionalTypes[MapKeys{caller, tstack.Val.(string)}]
		if ok {
			if base.IsMatchOptionalTypeForNumber(optional_types) {
				return false, nil
			}

			return false, b.invalidArgumentErrorMsg(
				tstack.Val,
				base.OptionalTypeToString(optional_types),
				"int or float.",
			)
		}

		if tstack.IsQuoted == false {
			return true, nil
		}

		*row = base.InformationWhenParsing[tstack.Val].Row

		return false, b.invalidArgumentErrorMsg(
			tstack.Val,
			"quoted symbol",
			"int or float",
		)

	default:
		*row = base.InformationWhenParsing[tstack.Val].Row

		return false, b.invalidArgumentErrorMsg(
			tstack.Val,
			base.TypeToString(tstack.Type),
			"int or float",
		)
	}
}

func (b BuiltinCompiler) isSymbolOrQuotedSymbol(
	tstack *base.S,
	row *int,
	caller string,
) (bool, error) {

	switch tstack.Type {
	case base.QUOTED_SYMBOL, base.ANY:
		return false, nil

	case base.SYMBOL:
		if tstack.IsQuoted == false {
			return true, nil
		}

		return false, nil

	default:
		*row = base.InformationWhenParsing[tstack.Val].Row

		return false, b.invalidArgumentErrorMsg(
			tstack.Val,
			base.TypeToString(tstack.Type),
			"quoted symbol",
		)
	}
}

func (b *BuiltinCompiler) invalidArgumentErrorMsg(
	error_value any,
	current_type_string string,
	want_type_string string,
) error {

	return fmt.Errorf(
		"%v is invalid argument type %v want %v",
		error_value,
		current_type_string,
		want_type_string,
	)
}
