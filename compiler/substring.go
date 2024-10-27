package compiler

import (
	"disco/base"
)

type SubString struct {
	BuiltinCompiler
}

func NewSubString() BuiltinCompilerIF {
	return &SubString{
		BuiltinCompiler{
			name:               "subseq",
			returnType:         base.STRING,
			firstArgumentType:  base.STRING,
			secondArgumentType: base.INT,
			thirdArgumentType:  base.INT,
			minArgumentCount:   3,
			maxArgumentCount:   3,
		},
	}
}

func init() {
	bc := NewSubString()
	BuiltinCompilers[bc.getName()] = bc
}

func (su *SubString) builtinCompile(
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

	codes = codeAppend(codes, base.SUBSTRING, caller, file_name, row)

	err = su.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (su *SubString) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := su.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	su.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopMultiStack(3)

	is_symbol, err :=
		su.isSymbolOrWantType(tstack[0], su.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			su.setFunctionArgumentTypes(tstack[0], caller, su.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	is_symbol, err =
		su.isSymbolOrWantType(tstack[1], su.secondArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			su.setFunctionArgumentTypes(tstack[1], caller, su.secondArgumentType, row)
		if err != nil {
			return err
		}
	}

	is_symbol, err =
		su.isSymbolOrWantType(tstack[2], su.thirdArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			su.setFunctionArgumentTypes(tstack[2], caller, su.thirdArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], su.getName())

	return nil
}
