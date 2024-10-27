package compiler

import (
	"disco/base"
)

type Remainder struct {
	BuiltinCompiler
}

func NewRemainder() BuiltinCompilerIF {
	return &Remainder{
		BuiltinCompiler{
			name:              "%",
			returnType:        base.INT,
			firstArgumentType: base.INT,
			minArgumentCount:  1,
			maxArgumentCount:  1000,
		},
	}
}

func init() {
	bc := NewRemainder()
	BuiltinCompilers[bc.getName()] = bc
}

func (r *Remainder) builtinCompile(
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

	codes = append(codes, base.REMAINDER)
	codes = append(codes, sLength(s))

	err = r.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (r *Remainder) typePropagation(
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

	tstack := TypeEnv.PopMultiStack(sLength(s))

	for _, t := range tstack {
		is_symbol, err :=
			r.isSymbolOrWantType(t, r.firstArgumentType, row, caller)
		if err != nil {
			return err
		}

		if is_symbol {
			err := r.setFunctionArgumentTypes(t, caller, r.firstArgumentType, row)
			if err != nil {
				return err
			}
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], r.getName())

	return nil
}
