package compiler

import (
	"disco/base"
)

type Print struct {
	BuiltinCompiler
}

func NewPrint() BuiltinCompilerIF {
	return &Print{
		BuiltinCompiler{
			name:             "print",
			returnType:       base.ANY,
			minArgumentCount: 1,
			maxArgumentCount: 1,
		},
	}
}

func init() {
	bc := NewPrint()
	BuiltinCompilers[bc.getName()] = bc

	alias := NewPrint()
	alias.setAlias("p")
	BuiltinCompilers[alias.getName()] = alias
}

func (pr *Print) builtinCompile(
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

	codes = codeAppend(codes, base.PRINT, caller, file_name, row)

	err = pr.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (pr *Print) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := pr.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	TypeEnv.PopStack()

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], pr.getName())

	return nil
}
