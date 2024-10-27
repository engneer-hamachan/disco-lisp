package compiler

import (
	"disco/base"
)

type Funcall struct {
	BuiltinCompiler
}

func NewFuncall() BuiltinCompilerIF {
	return &Funcall{
		BuiltinCompiler{
			name:              "funcall",
			returnType:        base.ANY,
			firstArgumentType: base.EXEC_FUNC,
			minArgumentCount:  1,
			maxArgumentCount:  1000,
		},
	}
}

func init() {
	bc := NewFuncall()
	BuiltinCompilers[bc.getName()] = bc
}

func (f *Funcall) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	codes, err :=
		compiler.recurisonCompile(codes, s.GetCdr(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes, err = compiler.Compile(codes, s.GetCar(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.FUNCALL, caller, file_name, row)
	codes = codeAppend(codes, sLength(s.GetCdr()), caller, file_name, row)

	err = f.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (f *Funcall) decideFname(s *base.S) string {
	var fname string

	switch s.GetCadar().Val.(string) {
	case "list-data":
		if s.GetCadar().GetCar().Val.(string) == "lambda" ||
			s.GetCadar().GetCar().Val.(string) == "fn" {

			fname = "peek-lambda"
		}
	default:
		fname = s.GetCadar().Val.(string)
	}

	return fname
}

func (f *Funcall) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	delete(FunctionOptionalTypes, "funcall")
	delete(FunctionReturnTypes, "funcall")

	err := f.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	f.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopStack()

	is_symbol, err :=
		f.isSymbolOrWantType(tstack, f.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := f.setFunctionArgumentTypes(tstack, caller, f.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	fname := f.decideFname(s)

	bf, ok := BuiltinCompilers[fname]
	if ok {
		err := bf.typePropagation(s.GetCdr(), caller, file_name, row)
		if err != nil {
			return err
		}
	}

	defin_function, ok := base.Globals[fname]
	if ok {
		bc := BuiltinCompiler{
			name:             fname,
			minArgumentCount: sLengthWithoutOptional(defin_function.GetCar()),
			maxArgumentCount: sLengthWithOptional(defin_function.GetCar()),
		}

		err = bc.checkArgumentCount(s.GetCdr(), row)
		if err != nil {
			return err
		}

		err =
			typePropagation(s.GetCdr(), defin_function.GetCar(), caller, fname, row)
		if err != nil {
			return err
		}
	}

	optional_types, ok := FunctionOptionalTypes[fname]
	if ok {
		FunctionOptionalTypes["funcall"] = optional_types
		FunctionReturnTypes["funcall"] = base.OPTIONAL
	}

	TypeEnv.PopStack()
	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], f.getName())

	return nil
}
