package compiler

import (
	"disco/base"
)

type Random struct {
	BuiltinCompiler
}

func NewRandom() BuiltinCompilerIF {
	return &Random{
		BuiltinCompiler{
			name:              "random",
			returnType:        base.INT,
			firstArgumentType: base.INT,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewRandom()
	BuiltinCompilers[bc.getName()] = bc
}

func (r *Random) builtinCompile(
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

	codes = codeAppend(codes, base.RANDOM, caller, file_name, row)

	err = r.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (r *Random) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := r.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	tstack := TypeEnv.PopStack()

	r.setFunctionReturnTypes(caller)

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
