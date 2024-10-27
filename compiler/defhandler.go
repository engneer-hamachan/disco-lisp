package compiler

import (
	"disco/base"
)

type DefHandler struct {
	BuiltinCompiler
}

func NewDefHandler() BuiltinCompilerIF {
	return &DefHandler{
		BuiltinCompiler{
			name:               "defhandler",
			returnType:         base.ANY,
			firstArgumentType:  base.STRING,
			secondArgumentType: base.EXEC_FUNC,
			thirdArgumentType:  base.LIST,
			minArgumentCount:   2,
			maxArgumentCount:   2,
		},
	}
}

func init() {
	bc := NewDefHandler()
	BuiltinCompilers[bc.getName()] = bc
}

func (d *DefHandler) builtinCompile(
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

	codes = codeAppend(codes, base.DEFHANDLER, caller, file_name, row)

	err = d.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (d *DefHandler) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := d.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	d.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopMultiStack(2)

	car := s.GetCar()
	cadr := s.GetCadr()

	is_symbol, err :=
		d.isSymbolOrWantType(car, d.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := d.setFunctionArgumentTypes(car, caller, d.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	is_symbol, err =
		d.isSymbolOrWantType(tstack[1], d.secondArgumentType, row, caller)
	if err != nil {
		*row = base.InformationWhenParsing[cadr.Val].Row

		return err
	}

	if is_symbol {
		err := d.setFunctionArgumentTypes(cadr, caller, d.thirdArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], d.getName())

	return nil
}
