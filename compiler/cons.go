package compiler

import (
	"disco/base"
)

type Cons struct {
	BuiltinCompiler
}

func NewCons() BuiltinCompilerIF {
	return &Cons{
		BuiltinCompiler{
			name:             "cons",
			returnType:       base.LIST,
			minArgumentCount: 2,
			maxArgumentCount: 2,
		},
	}
}

func init() {
	bc := NewCons()
	BuiltinCompilers[bc.getName()] = bc
}

func (c *Cons) builtinCompile(
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

	codes, err = compiler.Compile(codes, s.GetCadr(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.CONS, caller, file_name, row)

	err = c.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (c *Cons) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := c.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	c.setFunctionReturnTypes(caller)

	TypeEnv.PopMultiStack(2)

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], c.getName())

	return nil
}
