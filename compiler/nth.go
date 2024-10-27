package compiler

import (
	"disco/base"
)

type Nth struct {
	BuiltinCompiler
}

func NewNth() BuiltinCompilerIF {
	return &Nth{
		BuiltinCompiler{
			name:               "nth",
			returnType:         base.ANY,
			firstArgumentType:  base.INT,
			secondArgumentType: base.LIST,
			minArgumentCount:   2,
			maxArgumentCount:   2,
		},
	}
}

func init() {
	bc := NewNth()
	BuiltinCompilers[bc.getName()] = bc
}

func (n *Nth) builtinCompile(
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

	codes = append(codes, base.NTH)

	err = n.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (n *Nth) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := n.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	n.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopMultiStack(2)
	index := tstack[0]
	list_s := tstack[1]

	is_symbol, err :=
		n.isSymbolOrWantType(index, n.firstArgumentType, row, caller)

	if err != nil {
		return err
	}

	if is_symbol {
		err := n.setFunctionArgumentTypes(index, caller, n.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	is_symbol, err =
		n.isSymbolOrWantType(list_s, n.secondArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := n.setFunctionArgumentTypes(list_s, caller, n.secondArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], n.getName())

	return nil
}
