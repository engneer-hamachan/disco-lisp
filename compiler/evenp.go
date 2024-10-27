package compiler

import (
	"disco/base"
)

type Evenp struct {
	BuiltinCompiler
}

func NewEvenp() BuiltinCompilerIF {
	return &Evenp{
		BuiltinCompiler{
			name:              "even?",
			returnType:        base.ANY,
			firstArgumentType: base.INT,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewEvenp()
	BuiltinCompilers[bc.getName()] = bc
}

func (e *Evenp) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	codes, err := compiler.Compile(codes, s.GetCar(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.EVENP, caller, file_name, row)

	err = e.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (e *Evenp) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := e.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	e.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopStack()

	is_symbol, err :=
		e.isSymbolOrWantType(tstack, e.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := e.setFunctionArgumentTypes(tstack, caller, e.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], e.getName())

	return nil
}
