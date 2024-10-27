package compiler

import (
	"disco/base"
	"fmt"
)

type RunServer struct {
	BuiltinCompiler
}

func NewRunServer() BuiltinCompilerIF {
	return &RunServer{
		BuiltinCompiler{
			name:              "run-server",
			returnType:        base.ANY,
			firstArgumentType: base.LIST,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewRunServer()
	BuiltinCompilers[bc.getName()] = bc
}

func (r *RunServer) builtinCompile(
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

	codes = codeAppend(codes, base.RUNSERVER, caller, file_name, row)

	err = r.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (r *RunServer) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := r.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	r.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopStack()

	if tstack.Type != base.SYMBOL {
		*row = base.InformationWhenParsing[tstack.Val].Row

		return fmt.Errorf(
			"%v is invalid argument type %v want %v",
			tstack.Val,
			base.TypeToString(tstack.Type),
			"quote symbol",
		)
	}

	if s.GetCar().Type != base.LIST {
		err := r.setFunctionArgumentTypes(tstack, caller, r.firstArgumentType, row)
		if err != nil {
			return err
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], r.getName())

	return nil
}
