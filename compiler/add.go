package compiler

import (
	"disco/base"
)

type Add struct {
	BuiltinCompiler
}

func NewAdd() BuiltinCompilerIF {
	return &Add{
		BuiltinCompiler{
			name:              "+",
			returnType:        base.INT,
			firstArgumentType: base.INT,
			minArgumentCount:  2,
			maxArgumentCount:  1000,
		},
	}
}

func init() {
	bc := NewAdd()
	BuiltinCompilers[bc.getName()] = bc
}

func (a *Add) builtinCompile(
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

	codes = codeAppend(codes, base.ADD, caller, file_name, row)
	codes = codeAppend(codes, sLength(s), caller, file_name, row)

	err = a.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (a *Add) typePropagation(
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

	tstack := TypeEnv.PopMultiStack(sLength(s))

	for _, t := range tstack {
		is_symbol, err := a.isSymbolOrNumberType(t, row, caller)
		if err != nil {
			return err
		}

		if is_symbol {
			err := a.setFunctionArgumentTypes(t, caller, a.firstArgumentType, row)
			if err != nil {
				return err
			}
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], a.getName())

	return nil
}
