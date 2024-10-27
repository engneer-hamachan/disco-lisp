package compiler

import (
	"disco/base"
)

type Command struct {
	BuiltinCompiler
}

func NewCommand() BuiltinCompilerIF {
	return &Command{
		BuiltinCompiler{
			name:              "command",
			returnType:        base.STRING,
			firstArgumentType: base.STRING,
			minArgumentCount:  1,
			maxArgumentCount:  1000,
		},
	}
}

func init() {
	bc := NewCommand()
	BuiltinCompilers[bc.getName()] = bc
}

func (c *Command) builtinCompile(
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

	codes = codeAppend(codes, base.COMMAND, caller, file_name, row)
	codes = codeAppend(codes, sLength(s), caller, file_name, row)

	err = c.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (c *Command) typePropagation(
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

	tstack := TypeEnv.PopMultiStack(sLength(s))

	for _, t := range tstack {
		is_symbol, err :=
			c.isSymbolOrWantType(t, c.firstArgumentType, row, caller)
		if err != nil {
			return err
		}

		if is_symbol {
			err := c.setFunctionArgumentTypes(t, caller, c.firstArgumentType, row)
			if err != nil {
				return err
			}
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], c.getName())

	return nil
}
