package compiler

import (
	"disco/base"
	"fmt"
)

func compileFunction(
	chunks []any,
	s *base.S,
	args *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	fname := s.Val.(string)

	f, ok := base.Globals[fname]
	if !ok {
		*row = base.InformationWhenParsing[s.Val].Row
		return nil, fmt.Errorf("undefined function %s", fname)
	}

	bc := BuiltinCompiler{
		name:             fname,
		minArgumentCount: sLengthWithoutOptional(f.GetCar()),
		maxArgumentCount: sLengthWithOptional(f.GetCar()),
	}

	err := bc.checkArgumentCount(args, row)
	if err != nil {
		return nil, err
	}

	chunks, err = compiler.recurisonCompile(chunks, args, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	chunks = codeAppend(chunks, base.LDF, caller, file_name, row)
	chunks = codeAppend(chunks, fname, caller, file_name, row)
	chunks = codeAppend(chunks, base.CALL, caller, file_name, row)
	chunks = codeAppend(chunks, sLength(args), caller, file_name, row)

	err = typePropagation(args, f.GetCar(), caller, fname, row)
	if err != nil {
		return nil, err
	}

	return chunks, nil
}

func typePropagation(
	check_args *base.S,
	check_dargs *base.S,
	caller string,
	fname string,
	row *int,
) error {

	TypeEnv.PopMultiStack(sLength(check_args))

	FunctionReturnTypes[caller] = FunctionReturnTypes[fname]

	for check_args != nil && check_args.Type != base.NIL {
		car := check_args.GetCar()
		dcar := check_dargs.GetCar()

		define_argument_type, ok :=
			FunctionArgumentTypes[MapKeys{fname, dcar.Val.(string)}]
		if !ok || define_argument_type == base.ANY {
			check_args = check_args.GetCdr()
			check_dargs = check_dargs.GetCdr()

			continue
		}

		switch car.Type {
		case base.SYMBOL:
			err := checkSymbolType(car, define_argument_type, caller, row)
			if err != nil {
				return err
			}

			propagationSymbolType(car, dcar, caller, fname)

		case base.LIST:
			if isFunctionList(car) {
				err := checkFunctionReturnType(car, define_argument_type, row)
				if err != nil {
					return err
				}

				break
			}

			err := checkListType(car, define_argument_type, row)
			if err != nil {
				return err
			}

		case define_argument_type:
			break

		default:
			err := checkImmediateType(car, caller, define_argument_type, row)
			if err != nil {
				return err
			}
		}

		check_args = check_args.GetCdr()
		check_dargs = check_dargs.GetCdr()
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[fname], fname)

	return nil
}

func propagationSymbolType(
	car *base.S,
	dcar *base.S,
	caller string,
	fname string,
) {

	FunctionArgumentTypes[MapKeys{caller, car.Val.(string)}] =
		FunctionArgumentTypes[MapKeys{fname, dcar.Val.(string)}]
}

func checkFunctionReturnType(
	car *base.S,
	define_argument_type int,
	row *int,
) error {

	fname := car.GetCar().Val.(string)

	f_return_type, ok := FunctionReturnTypes[fname]
	if !ok {
		return nil
	}

	switch f_return_type {
	case define_argument_type, base.ANY:
		return nil

	case base.OPTIONAL:
		optional_types := FunctionOptionalTypes[fname]

		for _, types := range optional_types {
			if types == define_argument_type {
				return nil
			}
		}

		return fmt.Errorf(
			"function %s is return %s want %s",
			fname,
			base.OptionalTypeToString(optional_types),
			base.TypeToString(define_argument_type),
		)

	default:
		*row = base.InformationWhenParsing[fname].Row

		return fmt.Errorf(
			"function %s is return %s want %s",
			fname,
			base.TypeToString(f_return_type),
			base.TypeToString(define_argument_type),
		)
	}
}

func checkImmediateType(
	car *base.S,
	caller string,
	define_argument_type int,
	row *int,
) error {

	if define_argument_type == car.Type {
		return nil
	}

	*row = base.InformationWhenParsing[car.Val].Row

	return fmt.Errorf(
		"variable %v is %v want %v",
		car.Val,
		base.TypeToString(car.Type),
		base.TypeToString(define_argument_type),
	)
}

func checkSymbolType(
	car *base.S,
	define_argument_type int,
	caller string,
	row *int,
) error {

	if define_argument_type == base.ANY {
		return nil
	}

	if car.Type != base.SYMBOL {
		return nil
	}

	caller_argument_type, ok :=
		FunctionArgumentTypes[MapKeys{caller, car.Val.(string)}]
	if !ok {
		return nil
	}

	switch caller_argument_type {
	case define_argument_type, base.ANY:
		return nil

	case base.OPTIONAL:
		optional_types :=
			FunctionArgumentOptionalTypes[MapKeys{caller, car.Val.(string)}]

		if len(optional_types) == 0 ||
			base.IsMatchOptionalType(optional_types, define_argument_type) {

			return nil
		}

		return fmt.Errorf(
			"%v is invalid argument type %v want %v",
			car.Val,
			base.OptionalTypeToString(optional_types),
			base.TypeToString(define_argument_type),
		)

	default:
		*row = base.InformationWhenParsing[car.Val].Row

		return fmt.Errorf(
			"%v is %v want %v",
			car.Val,
			base.TypeToString(caller_argument_type),
			base.TypeToString(define_argument_type),
		)
	}
}

func isFunctionList(car *base.S) bool {
	if car.GetCar().Type == base.SYMBOL &&
		car.GetCar().Val.(string) != "quote" &&
		car.GetCar().Val.(string) != "function" {

		return true
	}

	return false
}

func isQuoteIntList(define_argument_type int, car *base.S) bool {
	if define_argument_type == base.INT &&
		car.GetCar().Val == "quote" &&
		car.GetCadr().Type == base.INT {

		return true
	}

	return false
}

func checkListType(car *base.S, define_argument_type int, row *int) error {
	if define_argument_type == base.EXEC_FUNC &&
		car.GetCar().Val == "function" {

		return nil
	}

	if define_argument_type == base.QUOTED_SYMBOL &&
		car.Type == base.LIST &&
		car.GetCar().Val.(string) == "quote" {

		return nil
	}

	if isQuoteIntList(define_argument_type, car) {
		return nil
	}

	if define_argument_type != base.LIST {
		*row = base.InformationWhenParsing[car.Val].Row

		return fmt.Errorf(
			"%v is %v want %v",
			car.Val,
			base.TypeToString(car.Type),
			base.TypeToString(define_argument_type),
		)
	}

	return nil
}
