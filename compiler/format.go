package compiler

import (
	"disco/base"
)

type Format struct {
	BuiltinCompiler
}

func NewFormat() BuiltinCompilerIF {
	return &Format{
		BuiltinCompiler{
			name:               "format",
			returnType:         base.STRING,
			firstArgumentType:  base.INT,
			secondArgumentType: base.STRING,
			minArgumentCount:   1,
			maxArgumentCount:   1000,
		},
	}
}

func init() {
	bc := NewFormat()
	BuiltinCompilers[bc.getName()] = bc
}

func (f *Format) builtinCompile(
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

	codes = codeAppend(codes, base.FORMAT, caller, file_name, row)
	codes = codeAppend(codes, sLength(s), caller, file_name, row)

	err = f.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (f *Format) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := f.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	f.setFunctionReturnTypes(caller)

	tstack := TypeEnv.PopMultiStack(sLength(s))

	args := tstack[1:]
	target := f.extractionReplaceTarget(tstack[0].Val.(string))

	for idx, t := range args {
		if len(target) < 1 {
			continue
		}

		switch target[idx] {
		case "a":
			is_symbol, err :=
				f.isSymbolOrWantType(t, f.secondArgumentType, row, caller)
			if err != nil {
				return err
			}

			if is_symbol {
				err := f.setFunctionArgumentTypes(t, caller, f.secondArgumentType, row)
				if err != nil {
					return err
				}
			}

		case "d":
			is_symbol, err := f.isSymbolOrNumberType(t, row, caller)
			if err != nil {
				return err
			}

			if is_symbol {
				err := f.setFunctionArgumentTypes(t, caller, f.firstArgumentType, row)
				if err != nil {
					return err
				}
			}
		}
	}

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], f.getName())

	return nil
}

func (f *Format) extractionReplaceTarget(str string) []string {
	var target []string

	for i := 0; i < len(str); i++ {
		if string(str[i]) == "~" && i+1 < len(str) {
			target = append(target, string(str[i+1]))
		}
	}

	return target
}
