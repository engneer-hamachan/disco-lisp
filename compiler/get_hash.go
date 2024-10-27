package compiler

import (
	"disco/base"
)

type GetHash struct {
	BuiltinCompiler
}

func NewGetHash() BuiltinCompilerIF {
	return &GetHash{
		BuiltinCompiler{
			name:               "gethash",
			returnType:         base.ANY,
			firstArgumentType:  base.HASH,
			secondArgumentType: base.QUOTED_SYMBOL,
			minArgumentCount:   2,
			maxArgumentCount:   2,
		},
	}
}

func init() {
	bc := NewGetHash()
	BuiltinCompilers[bc.getName()] = bc
}

func (g *GetHash) builtinCompile(
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

	codes = codeAppend(codes, base.GET_HASH, caller, file_name, row)

	err = g.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (g *GetHash) typePropagation(
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

	is_symbol, err :=
		g.isSymbolOrWantType(tstack[0], g.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			g.setFunctionArgumentTypes(tstack[0], caller, g.firstArgumentType, row)

		if err != nil {
			return err
		}
	}

	is_symbol, err = g.isSymbolOrQuotedSymbol(tstack[1], row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			g.setFunctionArgumentTypes(tstack[1], caller, g.secondArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], g.getName())

	return nil
}

func gethashCompile(s *base.S) *base.S {
	return base.Cons(
		base.MakeSym("gethash"),
		base.Cons(s.GetCar(),
			base.Cons(s.GetCadr(),
				base.NilAtom,
			),
		),
	)
}
