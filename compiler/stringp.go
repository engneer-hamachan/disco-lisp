package compiler

import (
	"disco/base"
)

type Stringp struct {
	BuiltinCompiler
}

func NewStringp() BuiltinCompilerIF {
	return &Stringp{
		BuiltinCompiler{
			name:             "str?",
			returnType:       base.ANY,
			minArgumentCount: 1,
			maxArgumentCount: 1,
		},
	}
}

func init() {
	bc := NewStringp()
	BuiltinCompilers[bc.getName()] = bc
}

func (st *Stringp) builtinCompile(
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

	codes = codeAppend(codes, base.STRINGP, caller, file_name, row)

	err = st.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (st *Stringp) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := st.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	st.setFunctionReturnTypes(caller)

	TypeEnv.PopStack()

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], st.getName())

	return nil
}
