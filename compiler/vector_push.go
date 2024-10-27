package compiler

import (
	"disco/base"
)

type VectorPush struct {
	BuiltinCompiler
}

func NewVectorPush() BuiltinCompilerIF {
	return &VectorPush{
		BuiltinCompiler{
			name:               "vector-push",
			returnType:         base.VECTOR,
			secondArgumentType: base.VECTOR,
			firstArgumentType:  base.ANY,
			minArgumentCount:   2,
			maxArgumentCount:   2,
		},
	}
}

func init() {
	bc := NewVectorPush()
	BuiltinCompilers[bc.getName()] = bc
}

func (v *VectorPush) builtinCompile(
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

	codes = codeAppend(codes, base.VECTOR_PUSH, caller, file_name, row)

	err = v.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (v *VectorPush) typePropagation(
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

	tstack := TypeEnv.PopMultiStack(2)

	is_symbol, err :=
		v.isSymbolOrWantType(tstack[0], v.secondArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			v.setFunctionArgumentTypes(tstack[0], caller, v.secondArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], v.getName())

	return nil
}
