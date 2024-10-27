package fever

import (
	"disco/base"
	"disco/compiler"
	"fmt"
	"os"
)

var VM = base.NewFeverMachine()

func Fever(
	codes []any,
	env *base.Environment,
	caller any,
) error {

	pc := 0
	size := len(codes) - 1
	var err error

	for pc <= size {
		executor := BuiltinExecutors[codes[pc].(int)]

		pc, err = executor.Execute(codes, pc, env, caller)
		if err != nil {
			var caller_string string

			if caller != nil {
				caller_string = caller.(string)
			}

			VM.Fatal(caller_string, err, pc)
			return fmt.Errorf("runtime error")
		}

		pc += 1
	}

	return nil
}

func makePosixArgs(os_args_idx int) *base.S {
	if os_args_idx >= len(os.Args) {
		return base.NilAtom
	}

	car := base.MakeString(os.Args[os_args_idx])
	os_args_idx++

	return base.Cons(car, makePosixArgs(os_args_idx))
}

func setPosixArgs() {
	if len(os.Args) < 3 {
		base.Globals["*argv*"] = base.NilAtom
		return
	}

	base.Globals["*argv*"] = makePosixArgs(2)
}

func FeverPreparaion() {
	//clear memory. when use compile.
	compiler.FunctionReturnTypes = make(map[string]int)
	compiler.FunctionArgumentTypes = make(map[compiler.MapKeys]int)
	compiler.FunctionOptionalTypes = make(map[string][]int)
	compiler.FunctionArgumentOptionalTypes = make(map[compiler.MapKeys][]int)

	setPosixArgs()
}
