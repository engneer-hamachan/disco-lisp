package compiler

import (
	"disco/base"
)

type Progn struct {
	BuiltinCompiler
}

func NewProgn() BuiltinCompilerIF {
	return &Progn{
		BuiltinCompiler{
			name:             "progn",
			returnType:       base.ANY,
			minArgumentCount: 1,
			maxArgumentCount: 1000,
		},
	}
}

func init() {
	bc := NewProgn()
	BuiltinCompilers[bc.getName()] = bc
}

func (pr *Progn) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	codes = codeAppend(codes, base.PROGN_START, caller, file_name, row)

	codes, err := compiler.recurisonCompile(codes, s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.PROGN_END, caller, file_name, row)

	err = pr.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (pr *Progn) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := pr.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	TypeEnv.PopMultiStack(sLength(s))

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], pr.getName())

	return nil
}
