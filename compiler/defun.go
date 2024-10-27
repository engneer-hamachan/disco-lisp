package compiler

import (
	"disco/base"
)

type DeFun struct {
	BuiltinCompiler
}

func NewDeFun() BuiltinCompilerIF {
	return &DeFun{}
}

func init() {
	BuiltinCompilers["def"] = NewDeFun()
}

func (d *DeFun) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	fname := s.GetCar().Val.(string)
	body := s.GetCddr()
	args := s.GetCadr()

	f := base.MakeFunc(fname, args)
	base.Globals[fname] = f

	d.setFunctionReturnTypes(caller)

	var fargs []*base.S

	for {
		if args.Type == base.NIL {
			break
		}

		if args.GetCar().Val.(string) == "&optional" {
			args = args.GetCdr()
			continue
		}

		fargs = append(fargs, args.GetCar())

		FunctionArgumentTypes[MapKeys{fname, args.GetCar().Val.(string)}] = base.ANY

		args = args.GetCdr()
	}

	var fbody []any

	fbody, err := compiler.recurisonCompile(fbody, body, fname, file_name, row)
	if err != nil {
		return nil, err
	}

	function_object := &base.F{
		Name: fname,
		Body: fbody,
		Args: fargs,
		Env:  nil,
	}

	codes = codeAppend(codes, base.DEFUN, caller, file_name, row)
	codes = codeAppend(codes, function_object, caller, file_name, row)

	err = d.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (d *DeFun) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], d.getName())

	return nil
}
