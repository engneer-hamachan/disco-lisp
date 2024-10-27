package compiler

import (
	"disco/base"
)

type Assoc struct {
	BuiltinCompiler
}

func NewAssoc() BuiltinCompilerIF {
	return &Assoc{
		BuiltinCompiler{
			name:               "assoc",
			returnType:         base.LIST,
			firstArgumentType:  base.ANY,
			secondArgumentType: base.LIST,
			minArgumentCount:   2,
			maxArgumentCount:   2,
		},
	}
}

func init() {
	bc := NewAssoc()
	BuiltinCompilers[bc.getName()] = bc
}

func (a *Assoc) builtinCompile(
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

	codes = codeAppend(codes, base.ASSOC, caller, file_name, row)

	err = a.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (a *Assoc) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := a.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	a.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopMultiStack(2)
	cadr := tstack[1]

	is_symbol, err :=
		a.isSymbolOrWantType(cadr, a.secondArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err = a.setFunctionArgumentTypes(cadr, caller, a.secondArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], a.getName())

	return nil
}
