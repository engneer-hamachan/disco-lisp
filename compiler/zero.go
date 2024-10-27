package compiler

import (
	"disco/base"
)

type Zerop struct {
	BuiltinCompiler
}

func NewZerop() BuiltinCompilerIF {
	return &Zerop{
		BuiltinCompiler{
			name:              "zero?",
			returnType:        base.ANY,
			firstArgumentType: base.INT,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewZerop()
	BuiltinCompilers[bc.getName()] = bc
}

func (z *Zerop) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	codes, err :=
		compiler.Compile(codes, base.MakeInt(int64(0)), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes, err = compiler.Compile(codes, s.GetCar(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.EQ, caller, file_name, row)

	err = z.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (z *Zerop) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := z.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	z.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopStack()

	is_symbol, err :=
		z.isSymbolOrWantType(tstack, z.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := z.setFunctionArgumentTypes(tstack, caller, z.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], z.getName())

	return nil
}
