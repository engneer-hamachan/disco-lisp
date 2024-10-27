package compiler

import (
	"disco/base"
)

type Function struct {
	BuiltinCompiler
}

func NewFunction() BuiltinCompilerIF {
	return &Function{
		BuiltinCompiler{
			name:              "function",
			returnType:        base.EXEC_FUNC,
			firstArgumentType: base.ANY,
			minArgumentCount:  1,
			maxArgumentCount:  1,
		},
	}
}

func init() {
	bc := NewFunction()
	BuiltinCompilers[bc.getName()] = bc
}

func (f *Function) builtinCompile(
	codes []any,
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	bf, ok := BuiltinCompilers[s.GetCar().Val.(string)]
	if ok {
		base.InformationWhenParsing["lambda"] =
			base.Info{
				FileName: *file_name,
				Row:      *row,
			}

		switch bf.getMinArgumentCount() {
		case 1:
			s = f.wrapInSingleAragumentLambda(s.GetCar().Val.(string))
		default:
			s = f.wrapInManyAragumentLambda(s.GetCar().Val.(string))
		}
	}

	codes, err := compiler.Compile(codes, s.GetCar(), caller, file_name, row)
	if err != nil {
		return nil, err
	}

	codes = codeAppend(codes, base.FUNCTION, caller, file_name, row)

	err = f.typePropagation(s, caller, file_name, row)
	if err != nil {
		return nil, err
	}

	return codes, nil
}

func (f *Function) wrapInSingleAragumentLambda(fname string) *base.S {
	s := base.Cons(
		base.Cons(
			base.MakeSym("fn"),
			base.Cons(
				base.Cons(
					base.MakeSym("x"),
					base.NilAtom,
				),
				base.Cons(
					base.Cons(
						base.MakeSym(fname),
						base.Cons(
							base.MakeSym("x"),
							base.NilAtom,
						),
					),
					base.NilAtom,
				),
			),
		),
		base.NilAtom,
	)

	return s
}

func (f *Function) wrapInManyAragumentLambda(fname string) *base.S {
	s := base.Cons(
		base.Cons(
			base.MakeSym("fn"),
			base.Cons(
				base.Cons(
					base.MakeSym("x"),
					base.Cons(
						base.MakeSym("y"),
						base.NilAtom,
					),
				),
				base.Cons(
					base.Cons(
						base.MakeSym(fname),
						base.Cons(
							base.MakeSym("x"),
							base.Cons(
								base.MakeSym("y"),
								base.NilAtom,
							),
						),
					),
					base.NilAtom,
				),
			),
		),
		base.NilAtom,
	)

	return s
}

func (f *Function) typePropagation(
	s *base.S,
	caller string,
	file_name *string,
	row *int,
) error {

	err := f.checkArgumentCount(s, row)
	if err != nil {
		return err
	}

	TypeEnv.PopMultiStack(sLength(s))

	f.setFunctionReturnTypes(caller)

	TypeEnv.PushDummyStack(FunctionReturnTypes[caller], f.getName())

	return nil
}
