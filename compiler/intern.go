package compiler

import (
	"disco/base"
)

type Intern struct {
	BuiltinCompiler
}

func NewIntern() BuiltinCompilerIF {
	return &Intern{
		BuiltinCompiler{
			name:              "intern",
			returnType:        base.QUOTED_SYMBOL,
			firstArgumentType: base.STRING,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewIntern()
	BuiltinCompilers[bc.getName()] = bc
}

func (i *Intern) builtinCompile(
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

	codes = codeAppend(codes, base.INTERN, caller, file_name, row)

	err = i.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (i *Intern) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	tstack := TypeEnv.PopStack()

	i.setFunctionReturnTypes(caller)

	is_symbol, err :=
		i.isSymbolOrWantType(tstack, i.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := i.setFunctionArgumentTypes(tstack, caller, i.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], i.getName())

	return nil
}
