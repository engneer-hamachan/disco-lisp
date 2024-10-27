package compiler

import (
	"disco/base"
)

type DefMacro struct {
	BuiltinCompiler
}

func NewDefMacro() BuiltinCompilerIF {
	return &DefMacro{}
}

func init() {
	BuiltinCompilers["mac"] = NewDefMacro()
}

func (d *DefMacro) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	macro_name := s.GetCar()
	args := s.GetCadr()
	body := s.GetCddr()

	macro := base.MakeMacro(macro_name.Val.(string), args, body)

	base.Globals[macro_name.Val.(string)] = macro

	err := d.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (d *DefMacro) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], "defmacro")

	return nil
}
