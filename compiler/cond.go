package compiler

import (
	"disco/base"
)

type Cond struct {
	BuiltinCompiler
}

func NewCond() BuiltinCompilerIF {
	return &Cond{
		BuiltinCompiler{
			name:             "cond",
			returnType:       base.ANY,
			minArgumentCount: 1,
			maxArgumentCount: 1000,
		},
	}
}

func init() {
	bc := NewCond()
	BuiltinCompilers[bc.getName()] = bc
}

func (c *Cond) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	var err error

	before_return_type := base.ANY

	for {
		codes, err = compiler.Compile(codes, s.GetCar(), caller, file_name, row)
		if err != nil {
			return nil, err
		}

		codes = codeAppend(codes, base.IF, caller, file_name, row)

		var then_codes []any

		then_codes, err =
			compiler.Compile(then_codes, s.GetCadr(), caller, file_name, row)
		if err != nil {
			return nil, err
		}

		block_return_type := FunctionReturnTypes[caller]

		c.typePropagationForOptionalType(
			before_return_type,
			block_return_type,
			caller,
		)

		before_return_type = FunctionReturnTypes[caller]

		codes = codeAppend(codes, len(then_codes)+2, caller, file_name, row)

		for _, code := range then_codes {
			codes = codeAppend(codes, code, caller, file_name, row)
		}

		codes = codeAppend(codes, base.LJMP, caller, file_name, row)
		codes = codeAppend(codes, "COND_END", caller, file_name, row)

		if s.GetCddr().Type != base.NIL {
			s = s.GetCddr()

			continue
		}

		break
	}

	codes = codeAppend(codes, base.LD, caller, file_name, row)
	codes = codeAppend(codes, base.NilAtom, caller, file_name, row)
	codes = codeAppend(codes, base.LABEL, caller, file_name, row)
	codes = codeAppend(codes, "COND_END", caller, file_name, row)

	err = c.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (c *Cond) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := c.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	err = c.checkArgumentCountIsEven(s, row)
	if err != nil {
		return err
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], c.getName())

	return nil
}

func (c *Cond) typePropagationForOptionalType(
	before_return_type int,
	block_return_type int,
	caller string,
) {

	TypeEnv.PopMultiStack(2)

	switch before_return_type {
	case block_return_type, base.ANY:
		FunctionReturnTypes[caller] = block_return_type

	case base.OPTIONAL:
		FunctionReturnTypes[caller] = base.OPTIONAL

		if !base.IsMatchOptionalType(
			FunctionOptionalTypes[caller],
			block_return_type,
		) {

			c.appendFunctionOptionalTypes(caller, block_return_type)
		}

	default:
		c.clearFunctionOptionalTypes(caller)
		c.setFunctionChoiceReturnTypes(caller, base.OPTIONAL)
		c.appendFunctionOptionalTypes(caller, before_return_type)
		c.appendFunctionOptionalTypes(caller, block_return_type)
	}
}
