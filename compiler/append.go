package compiler

import (
	"disco/base"
	"disco/predicater"
)

type Append struct {
	BuiltinCompiler
}

func NewAppend() BuiltinCompilerIF {
	return &Append{
		BuiltinCompiler{
			name:               "append",
			returnType:         base.LIST,
			firstArgumentType:  base.LIST,
			secondArgumentType: base.ANY,
			minArgumentCount:   2,
			maxArgumentCount:   2,
		},
	}
}

func init() {
	bc := NewAppend()
	BuiltinCompilers[bc.getName()] = bc
}

func (a *Append) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	codes, err := compiler.recurisonCompile(codes, s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.APPEND, caller, file_name, row)

	err = a.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (a *Append) Append(car *base.S, cdr *base.S) *base.S {
	if predicater.Nilp(car) {
		return cdr
	}

	return base.Cons(car.GetCar(), a.Append(car.GetCdr(), cdr))
}

func (a *Append) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := a.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	a.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopMultiStack(2)

	is_symbol, err :=
		a.isSymbolOrWantType(tstack[0], a.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			a.setFunctionArgumentTypes(tstack[0], caller, a.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	is_symbol, err =
		a.isSymbolOrWantType(tstack[1], a.secondArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			a.setFunctionArgumentTypes(tstack[1], caller, a.secondArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], a.getName())

	return nil
}
