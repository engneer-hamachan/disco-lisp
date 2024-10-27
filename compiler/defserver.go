package compiler

import (
	"disco/base"
)

type DefServer struct {
	BuiltinCompiler
}

func NewDefServer() BuiltinCompilerIF {
	return &DefServer{
		BuiltinCompiler{
			name:               "defserver",
			returnType:         base.ANY,
			firstArgumentType:  base.QUOTED_SYMBOL,
			secondArgumentType: base.LIST,
			thirdArgumentType:  base.STRING,
			minArgumentCount:   2,
			maxArgumentCount:   2,
		},
	}
}

func init() {
	bc := NewDefServer()
	BuiltinCompilers[bc.getName()] = bc
}

func (d *DefServer) builtinCompile(
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

	codes = codeAppend(codes, base.DEFSERVER, caller, file_name, row)

	err = d.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (d *DefServer) typePropagation(
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

	_, err = d.isSymbolOrQuotedSymbol(tstack[0], row, caller)
	if err != nil {
		return err
	}

	if s.GetCar().Type != base.LIST {
		err :=
			d.setFunctionArgumentTypes(tstack[0], caller, d.secondArgumentType, row)

		if err != nil {
			return err
		}
	}

	is_symbol, err :=
		d.isSymbolOrWantType(tstack[1], d.thirdArgumentType, row, caller)
	if err != nil {
		*row = base.InformationWhenParsing[s.GetCadr().Val].Row

		return err
	}

	if is_symbol {
		err :=
			d.setFunctionArgumentTypes(tstack[1], caller, d.thirdArgumentType, row)

		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], d.getName())

	return nil
}
