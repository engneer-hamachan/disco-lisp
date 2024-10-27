package compiler

import (
	"disco/base"
)

type RegexpReplaceAll struct {
	BuiltinCompiler
}

func NewRegexpReplaceAll() BuiltinCompilerIF {
	return &RegexpReplaceAll{
		BuiltinCompiler{
			name:               "regexp-replace",
			returnType:         base.STRING,
			firstArgumentType:  base.STRING,
			secondArgumentType: base.STRING,
			thirdArgumentType:  base.STRING,
			minArgumentCount:   3,
			maxArgumentCount:   3,
		},
	}
}

func init() {
	bc := NewRegexpReplaceAll()
	BuiltinCompilers[bc.getName()] = bc
}

func (r *RegexpReplaceAll) builtinCompile(
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

	codes = codeAppend(codes, base.REGEXP_REPLACE_ALL, caller, file_name, row)

	err = r.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (r *RegexpReplaceAll) typePropagation(
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

	tstack := TypeEnv.PopMultiStack(3)

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
