package compiler

import (
	"disco/base"
	"fmt"
)

var TypeEnv = NewTypeEnvironment()

type TypeEnvironment struct {
	Stack []*base.S
}

func NewTypeEnvironment() TypeEnvironment {
	var stack []*base.S
	return TypeEnvironment{Stack: stack}
}

func (t *TypeEnvironment) size() int {
	return len(t.Stack)
}

func (t *TypeEnvironment) tailIdx() int {
	return t.size() - 1
}

func (t *TypeEnvironment) PushStack(s *base.S) {
	t.Stack = append(t.Stack, s)
}

func (t *TypeEnvironment) PushDummyStack(types int, value string) {
	dummy_s := base.MakeDummyS(types, value)
	t.PushStack(dummy_s)
}

func (t *TypeEnvironment) PopStack() *base.S {
	if len(t.Stack) > 0 {
		tstack := t.Stack[t.tailIdx()]
		t.Stack = t.Stack[:t.tailIdx()]
		return tstack
	}

	return nil
}

func (t *TypeEnvironment) PopMultiStack(
	length int,
) []*base.S {

	if len(t.Stack) < length {
		return []*base.S{}
	}

	part_of_stack := t.Stack[len(t.Stack)-length:]
	t.Stack = t.Stack[:len(t.Stack)-length]

	return part_of_stack
}

func (t *TypeEnvironment) DumpStack() {
	ct := 0

	for {
		if len(t.Stack)-1 < ct {
			break
		}

		fmt.Println(t.Stack[ct])
		ct++
	}
}
