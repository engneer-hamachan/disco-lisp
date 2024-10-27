package compiler

import (
	"disco/base"
)

type SetHash struct {
	BuiltinCompiler
}

func NewSethash() BuiltinCompilerIF {
	return &SetHash{
		BuiltinCompiler{
			name:               "sethash",
			returnType:         base.ANY,
			firstArgumentType:  base.QUOTED_SYMBOL,
			secondArgumentType: base.HASH,
			thirdArgumentType:  base.ANY,
			minArgumentCount:   3,
			maxArgumentCount:   3,
		},
	}
}

func init() {
	bc := NewSethash()
	BuiltinCompilers[bc.getName()] = bc
}

func (sh *SetHash) builtinCompile(
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

	codes = codeAppend(codes, base.SET_HASH, caller, file_name, row)

	err = sh.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (sh *SetHash) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := sh.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	tstack := TypeEnv.PopMultiStack(3)

	is_symbol, err := sh.isSymbolOrQuotedSymbol(tstack[0], row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			sh.setFunctionArgumentTypes(tstack[0], caller, sh.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	is_symbol, err =
		sh.isSymbolOrWantType(tstack[1], sh.secondArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			sh.setFunctionArgumentTypes(tstack[1], caller, sh.secondArgumentType, row)
		if err != nil {
			return err
		}
	}

	is_symbol, err =
		sh.isSymbolOrWantType(tstack[2], sh.thirdArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			sh.setFunctionArgumentTypes(tstack[2], caller, sh.thirdArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], sh.getName())

	return nil
}

func sethashCompile(s *base.S) *base.S {
	return base.Cons(
		base.MakeSym("sethash"),
		base.Cons(s.GetCadr(),
			base.Cons(s.GetCar(),
				base.Cons(s.GetCaddr(), base.NilAtom),
			),
		),
	)
}
