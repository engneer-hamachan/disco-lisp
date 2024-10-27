package compiler

import (
	"disco/base"
)

type VectorLen struct {
	BuiltinCompiler
}

func NewVectorLen() BuiltinCompilerIF {
	return &VectorLen{
		BuiltinCompiler{
			name:              "vector-len",
			returnType:        base.INT,
			firstArgumentType: base.VECTOR,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewVectorLen()
	BuiltinCompilers[bc.getName()] = bc
}

func (v *VectorLen) builtinCompile(
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

	codes = codeAppend(codes, base.VECTOR_LEN, caller, file_name, row)

	err = v.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (v *VectorLen) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := v.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	v.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopStack()

	is_symbol, err :=
		v.isSymbolOrWantType(tstack, v.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			v.setFunctionArgumentTypes(tstack, caller, v.firstArgumentType, row)

		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], v.getName())

	return nil
}
