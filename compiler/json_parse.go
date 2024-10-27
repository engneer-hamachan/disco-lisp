package compiler

import (
	"disco/base"
)

type JsonParse struct {
	BuiltinCompiler
}

func NewJsonParse() BuiltinCompilerIF {
	return &JsonParse{
		BuiltinCompiler{
			name:              "json-parse",
			returnType:        base.LIST,
			firstArgumentType: base.STRING,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewJsonParse()
	BuiltinCompilers[bc.getName()] = bc
}

func (j *JsonParse) builtinCompile(
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

	codes = codeAppend(codes, base.JSON_PARSE, caller, file_name, row)

	err = j.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (j *JsonParse) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := j.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	j.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopStack()

	is_symbol, err :=
		j.isSymbolOrWantType(tstack, j.firstArgumentType, row, caller)
	if err != nil {
		return err
	}

	if is_symbol {
		err := j.setFunctionArgumentTypes(tstack, caller, j.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], j.getName())

	return nil
}
