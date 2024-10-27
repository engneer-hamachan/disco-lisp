package compiler

import (
	"disco/base"
)

type Close struct {
	BuiltinCompiler
}

func NewClose() BuiltinCompilerIF {
	return &Close{
		BuiltinCompiler{
			name:              "close",
			returnType:        base.ANY,
			firstArgumentType: base.FP,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewClose()
	BuiltinCompilers[bc.getName()] = bc
}

func (c *Close) builtinCompile(
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

	codes = codeAppend(codes, base.CLOSE, caller, file_name, row)

	err = c.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (c *Close) typePropagation(
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

	tstack := TypeEnv.PopStack()

	is_symbol, err :=
		c.isSymbolOrWantType(tstack, c.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := c.setFunctionArgumentTypes(tstack, caller, c.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], c.getName())

	return nil
}
