package compiler

import (
	"disco/base"
)

type GetQuery struct {
	BuiltinCompiler
}

func NewGetQuery() BuiltinCompilerIF {
	return &GetQuery{
		BuiltinCompiler{
			name:              "get-query",
			returnType:        base.STRING,
			firstArgumentType: base.QUOTED_SYMBOL,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewGetQuery()
	BuiltinCompilers[bc.getName()] = bc
}

func (g *GetQuery) builtinCompile(
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

	codes = codeAppend(codes, base.GETQUERY, caller, file_name, row)

	err = g.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (g *GetQuery) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	g.setFunctionReturnTypes(caller)

	err := g.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	tstack := TypeEnv.PopStack()

	_, err = g.isSymbolOrQuotedSymbol(tstack, row, caller)
	if err != nil {
		return err
	}

	if s.GetCar().Type != base.LIST {
		err := g.setFunctionArgumentTypes(tstack, caller, g.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], g.getName())

	return nil
}
