package compiler

import (
	"disco/base"
)

type Ge struct {
	BuiltinCompiler
}

func NewGe() BuiltinCompilerIF {
	return &Ge{
		BuiltinCompiler{
			name:              ">=",
			returnType:        base.ANY,
			firstArgumentType: base.INT,
			minArgumentCount:  2,
			maxArgumentCount:  2,
		},
	}
}

func init() {
	bc := NewGe()
	BuiltinCompilers[bc.getName()] = bc
}

func (g *Ge) builtinCompile(
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

	codes = codeAppend(codes, base.GE, caller, file_name, row)

	err = g.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (g *Ge) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := g.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	g.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopMultiStack(2)

	for _, t := range tstack {
		is_symbol, err := g.isSymbolOrNumberType(t, row, caller)
		if err != nil {
			return err
		}

		if is_symbol {
			err := g.setFunctionArgumentTypes(t, caller, g.firstArgumentType, row)
			if err != nil {
				return err
			}
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], g.getName())

	return nil
}
