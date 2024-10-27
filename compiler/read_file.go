package compiler

import (
	"disco/base"
)

type ReadFile struct {
	BuiltinCompiler
}

func NewReadFile() BuiltinCompilerIF {
	return &ReadFile{
		BuiltinCompiler{
			name:              "read-file",
			returnType:        base.STRING,
			firstArgumentType: base.FP,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewReadFile()
	BuiltinCompilers[bc.getName()] = bc
}

func (r *ReadFile) builtinCompile(
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

	codes = codeAppend(codes, base.READ_FILE, caller, file_name, row)

	err = r.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (r *ReadFile) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := r.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	r.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopStack()

	is_symbol, err :=
		r.isSymbolOrWantType(tstack, r.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := r.setFunctionArgumentTypes(tstack, caller, r.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], r.getName())

	return nil
}
