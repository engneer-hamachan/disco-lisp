package compiler

import (
	"disco/base"
)

type Open struct {
	BuiltinCompiler
}

func NewOpen() BuiltinCompilerIF {
	return &Open{
		BuiltinCompiler{
			name:               "open",
			returnType:         base.FP,
			firstArgumentType:  base.STRING,
			secondArgumentType: base.STRING,
			minArgumentCount:   2,
			maxArgumentCount:   2,
		},
	}
}

func init() {
	bc := NewOpen()
	BuiltinCompilers[bc.getName()] = bc
}

func (o *Open) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	var err error

	codes, err = compiler.recurisonCompile(codes, s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.OPEN, caller, file_name, row)

	err = o.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (o *Open) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := o.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	tstack := TypeEnv.PopMultiStack(2)

	path := tstack[0]
	mode := tstack[1]

	is_symbol, err := o.isSymbolOrWantType(path, o.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := o.setFunctionArgumentTypes(path, caller, o.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	is_symbol, err = o.isSymbolOrWantType(mode, o.secondArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := o.setFunctionArgumentTypes(mode, caller, o.secondArgumentType, row)
		if err != nil {
			return err
		}
	}

	o.setFunctionReturnTypes(caller)

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], o.getName())

	return nil
}
