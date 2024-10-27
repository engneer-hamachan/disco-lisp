package fever

import (
	"disco/base"
	"fmt"
)

type Mapcar struct{}

func NewMapcar() BuiltinFeverIF {
	return &Mapcar{}
}

func init() {
	BuiltinExecutors[base.MAPCAR] = NewMapcar()
}

func (m *Mapcar) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	var f *base.F

	s := VM.PopStack().Val

	switch s.(type) {
	case *base.F:
		f = s.(*base.F)

	default:
		return pc, fmt.Errorf("not function args.")
	}

	f.Env.PushStack()

	var ct int

	stack_map := make(map[int]*base.S)
	part_of_stack := VM.PopMultiStack(codes[pc].(int))
	length_of_args := len(part_of_stack)

	for idx, s := range part_of_stack {
		stack_map[idx] = s
	}

	for stack_map[0].GetCar().Type != base.NIL {
		for idx, _ := range part_of_stack {
			VM.PushStack(stack_map[idx].GetCar())
			stack_map[idx] = stack_map[idx].GetCdr()
		}

		setArgs(f.Env, f.Args, length_of_args)
		Fever(f.Body, f.Env, f.Name)

		ct++
	}

	f.Env.PopStack()

	part_of_stack = VM.PopMultiStack(ct)
	VM.PushStack(m.consingResult(part_of_stack))

	return pc, nil
}

func (m *Mapcar) consingResult(part_of_stack []*base.S) *base.S {
	if len(part_of_stack) == 0 {
		return base.NilAtom
	}

	return base.Cons(part_of_stack[0], m.consingResult(part_of_stack[1:]))
}
