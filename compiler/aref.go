package compiler

import (
	"disco/base"
)

type Aref struct {
	BuiltinCompiler
}

func NewAref() BuiltinCompilerIF {
	return &Aref{
		BuiltinCompiler{
			name:               "aref",
			returnType:         base.ANY,
			firstArgumentType:  base.INT,
			secondArgumentType: base.VECTOR,
			minArgumentCount:   2,
			maxArgumentCount:   2,
		},
	}
}

func init() {
	bc := NewAref()
	BuiltinCompilers[bc.getName()] = bc
}

func (n *Aref) builtinCompile(
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

	codes = append(codes, base.AREF)

	err = n.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (n *Aref) typePropagation(
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
	vector_s := tstack[1]

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
		n.isSymbolOrWantType(vector_s, n.secondArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			n.setFunctionArgumentTypes(vector_s, caller, n.secondArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], n.getName())

	return nil
}

func arefCompile(s *base.S) *base.S {
	return base.Cons(
		base.MakeSym("aref"),
		base.Cons(s.GetCadr(),
			base.Cons(s.GetCar(),
				base.NilAtom,
			),
		),
	)
}
