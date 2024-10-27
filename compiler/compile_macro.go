package compiler

import (
	"disco/base"
	"disco/predicater"
)

func compileMacro(
	codes []any,
	f *base.S,
	args *base.S,
	caller string,
	file_name *string,
	row *int,
) ([]any, error) {

	dargs := f.GetCar()
	body := f.GetCadr()

	env := base.NewEnvironment()

	env.PushStack()

	err := setMacroArgs(&env, f, dargs, args)
	if err != nil {
		return nil, err
	}

	if predicater.HasName(body.GetCar(), "quasi-quote") {
		body = macroApply(&env, body.GetCadr())
	}

	env.PopStack()

	return compiler.Compile(codes, body, caller, file_name, row)
}

func setMacroArgs(
	env *base.Environment,
	f *base.S,
	dargs *base.S,
	args *base.S,
) error {

	if args.Type == base.NIL && dargs.Type == base.NIL {
		return nil
	}

	if dargs.GetCar().Val.(string) == "&rest" {
		env.SetSymbolValueForPeek(dargs.GetCadr(), args)
		return nil
	}

	env.SetSymbolValueForPeek(dargs.GetCar(), args.GetCar())

	err := setMacroArgs(env, f, dargs.GetCdr(), args.GetCdr())
	if err != nil {
		return err
	}

	return nil
}

func macroApply(env *base.Environment, body *base.S) *base.S {
	if predicater.Nilp(body) {
		return base.NilAtom
	}

	if predicater.Atomp(body) {
		return body
	}

	if predicater.Listp(body) &&
		predicater.HasName(body.GetCaar(), "unquote-splicing") {

		v, ok := env.GetSymbolValue(body.GetCadar())
		if ok {
			return base.Append(v, macroApply(env, body.GetCdr()))
		}

		return base.Append(body.GetCadar(), macroApply(env, body.GetCdr()))
	}

	if predicater.Listp(body) &&
		predicater.HasName(body.GetCaar(), "unquote") {

		v, ok := env.GetSymbolValue(body.GetCadar())
		if ok {
			return base.Cons(v, macroApply(env, body.GetCdr()))
		}

		return base.Cons(body.GetCadar(), macroApply(env, body.GetCdr()))
	}

	return base.Cons(
		macroApply(env, body.GetCar()),
		macroApply(env, body.GetCdr()),
	)
}
