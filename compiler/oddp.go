package compiler

import (
	"disco/base"
)

type Oddp struct {
	BuiltinCompiler
}

func NewOddp() BuiltinCompilerIF {
	return &Oddp{
		BuiltinCompiler{
			name:              "odd?",
			returnType:        base.ANY,
			firstArgumentType: base.INT,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewOddp()
	BuiltinCompilers[bc.getName()] = bc
}

func (o *Oddp) builtinCompile(
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

	codes = codeAppend(codes, base.ODDP, caller, file_name, row)

	err = o.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (o *Oddp) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := o.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	o.setFunctionReturnTypes(caller)

	args := TypeEnv.PopStack()

	is_symbol, err :=
		o.isSymbolOrWantType(args, o.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := o.setFunctionArgumentTypes(args, caller, o.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], o.getName())

	return nil
}
