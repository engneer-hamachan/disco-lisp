package compiler

import (
	"disco/base"
)

type Write struct {
	BuiltinCompiler
}

func NewWrite() BuiltinCompilerIF {
	return &Write{
		BuiltinCompiler{
			name:               "write",
			returnType:         base.ANY,
			firstArgumentType:  base.FP,
			secondArgumentType: base.STRING,
			minArgumentCount:   2,
			maxArgumentCount:   2,
		},
	}
}

func init() {
	bc := NewWrite()
	BuiltinCompilers[bc.getName()] = bc
}

func (w *Write) builtinCompile(
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

	codes = codeAppend(codes, base.WRITE, caller, file_name, row)

	err = w.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (w *Write) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := w.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	tstack := TypeEnv.PopMultiStack(2)

	is_symbol, err :=
		w.isSymbolOrWantType(tstack[0], w.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			w.setFunctionArgumentTypes(tstack[0], caller, w.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	is_symbol, err =
		w.isSymbolOrWantType(tstack[1], w.secondArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err :=
			w.setFunctionArgumentTypes(tstack[1], caller, w.secondArgumentType, row)
		if err != nil {
			return err
		}
	}

	w.setFunctionReturnTypes(caller)

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], w.getName())

	return nil
}
