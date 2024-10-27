package compiler

import (
	"disco/base"
)

type Split struct {
	BuiltinCompiler
}

func NewSplit() BuiltinCompilerIF {
	return &Split{
		BuiltinCompiler{
			name:               "split",
			returnType:         base.LIST,
			firstArgumentType:  base.STRING,
			secondArgumentType: base.STRING,
			minArgumentCount:   2,
			maxArgumentCount:   2,
		},
	}
}

func init() {
	bc := NewSplit()
	BuiltinCompilers[bc.getName()] = bc
}

func (sp *Split) builtinCompile(
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

	codes = codeAppend(codes, base.SPLIT, caller, file_name, row)

	err = sp.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (sp *Split) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := sp.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	sp.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopMultiStack(2)

	is_symbol, err :=
		sp.isSymbolOrWantType(tstack[0], sp.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			sp.setFunctionArgumentTypes(tstack[0], caller, sp.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	is_symbol, err =
		sp.isSymbolOrWantType(tstack[1], sp.secondArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			sp.setFunctionArgumentTypes(tstack[1], caller, sp.secondArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], sp.getName())

	return nil
}
