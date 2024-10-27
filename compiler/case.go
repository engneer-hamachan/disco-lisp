package compiler

import (
	"disco/base"
)

type Case struct {
	BuiltinCompiler
}

func NewCase() BuiltinCompilerIF {
	return &Case{
		BuiltinCompiler{
			name:              "case",
			returnType:        base.ANY,
			firstArgumentType: base.ANY,
			minArgumentCount:  1,
			maxArgumentCount:  1000,
		},
	}
}

func init() {
	bc := NewCase()
	BuiltinCompilers[bc.getName()] = bc
}

func (c *Case) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	var err error

	test_s := s.GetCar()
	s = s.GetCdr()

	before_return_type := base.ANY

	for {
		codes, err = compiler.Compile(codes, test_s, caller, file_name, row)
		if err != nil {
			return nil, err
		}

		codes, err = compiler.Compile(codes, s.GetCaar(), caller, file_name, row)
		if err != nil {
			return nil, err
		}

		codes = codeAppend(codes, base.CASE, caller, file_name, row)

		var then_codes []any

		then_codes, err =
			compiler.Compile(then_codes, s.GetCadar(), caller, file_name, row)

		if err != nil {
			return nil, err
		}

		then_return_type := FunctionReturnTypes[caller]

		c.typePropagationForOptionalType(
			before_return_type,
			then_return_type,
			caller,
		)

		before_return_type = then_return_type

		codes = codeAppend(codes, len(then_codes)+2, caller, file_name, row)

		for _, code := range then_codes {
			codes = codeAppend(codes, code, caller, file_name, row)
		}

		codes = codeAppend(codes, base.LJMP, caller, file_name, row)
		codes = codeAppend(codes, "CASE_END", caller, file_name, row)

		if s.GetCdr().Type != base.NIL {
			s = s.GetCdr()

			continue
		}

		break
	}

	codes = codeAppend(codes, base.LD, caller, file_name, row)
	codes = codeAppend(codes, base.NilAtom, caller, file_name, row)
	codes = codeAppend(codes, base.LABEL, caller, file_name, row)
	codes = codeAppend(codes, "CASE_END", caller, file_name, row)

	err = c.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (c *Case) typePropagationForOptionalType(
	before_return_type int,
	then_return_type int,
	caller string,
) {

	TypeEnv.PopMultiStack(3)

	switch before_return_type {
	case then_return_type, base.ANY:
		FunctionReturnTypes[caller] = then_return_type

	case base.OPTIONAL:
		c.setFunctionChoiceReturnTypes(caller, base.OPTIONAL)

		if !base.IsMatchOptionalType(
			FunctionOptionalTypes[caller],
			then_return_type,
		) {

			c.appendFunctionOptionalTypes(caller, then_return_type)
		}

	default:
		c.clearFunctionOptionalTypes(caller)
		c.setFunctionChoiceReturnTypes(caller, base.OPTIONAL)
		c.appendFunctionOptionalTypes(caller, before_return_type)
		c.appendFunctionOptionalTypes(caller, then_return_type)
	}
}

func (c *Case) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := c.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], c.getName())

	return nil
}
