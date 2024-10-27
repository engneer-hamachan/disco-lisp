package compiler

import (
	"disco/base"
)

type Le struct {
	BuiltinCompiler
}

func NewLe() BuiltinCompilerIF {
	return &Le{
		BuiltinCompiler{
			name:              "<=",
			returnType:        base.ANY,
			firstArgumentType: base.INT,
			minArgumentCount:  2,
			maxArgumentCount:  2,
		},
	}
}

func init() {
	bc := NewLe()
	BuiltinCompilers[bc.getName()] = bc
}

func (l *Le) builtinCompile(
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

	codes = codeAppend(codes, base.LE, caller, file_name, row)

	err = l.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (l *Le) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := l.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	l.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopMultiStack(2)

	for _, t := range tstack {
		is_symbol, err := l.isSymbolOrNumberType(t, row, caller)
		if err != nil {
			return err
		}

		if is_symbol {
			err := l.setFunctionArgumentTypes(t, caller, l.firstArgumentType, row)
			if err != nil {
				return err
			}
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], l.getName())

	return nil
}
