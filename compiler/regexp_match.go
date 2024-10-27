package compiler

import (
	"disco/base"
)

type RegexpMatch struct {
	BuiltinCompiler
}

func NewRegexpMatch() BuiltinCompilerIF {
	return &RegexpMatch{
		BuiltinCompiler{
			name:               "regexp-match",
			returnType:         base.STRING,
			firstArgumentType:  base.STRING,
			secondArgumentType: base.STRING,
			thirdArgumentType:  base.STRING,
			minArgumentCount:   2,
			maxArgumentCount:   2,
		},
	}
}

func init() {
	bc := NewRegexpMatch()
	BuiltinCompilers[bc.getName()] = bc
}

func (r *RegexpMatch) builtinCompile(
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

	codes = codeAppend(codes, base.REGEXP_MATCH, caller, file_name, row)

	err = r.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (r *RegexpMatch) typePropagation(
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

	tstack := TypeEnv.PopMultiStack(2)

	for _, arg := range tstack {
		is_symbol, err :=
			r.isSymbolOrWantType(arg, r.firstArgumentType, row, caller)
		if err != nil {
			return err
		}

		if is_symbol {
			err := r.setFunctionArgumentTypes(arg, caller, r.firstArgumentType, row)
			if err != nil {
				return err
			}
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], r.getName())

	return nil
}
