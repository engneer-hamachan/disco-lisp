package compiler

import (
	"disco/base"
)

type Div struct {
	BuiltinCompiler
}

func NewDiv() BuiltinCompilerIF {
	return &Div{
		BuiltinCompiler{
			name:              "/",
			returnType:        base.INT,
			firstArgumentType: base.INT,
			minArgumentCount:  2,
			maxArgumentCount:  1000,
		},
	}
}

func init() {
	bc := NewDiv()
	BuiltinCompilers[bc.getName()] = bc
}

func (d *Div) builtinCompile(
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

	codes = append(codes, base.DIV)
	codes = append(codes, sLength(s))

	err = d.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (d *Div) typePropagation(
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

	tstack := TypeEnv.PopMultiStack(sLength(s))

	for _, t := range tstack {
		is_symbol, err := d.isSymbolOrNumberType(t, row, caller)
		if err != nil {
			return err
		}

		if is_symbol {
			err := d.setFunctionArgumentTypes(t, caller, d.firstArgumentType, row)
			if err != nil {
				return err
			}
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], d.getName())

	return nil
}
