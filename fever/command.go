package fever

import (
	"disco/base"
	"os/exec"
)

type Command struct{}

func NewCommand() BuiltinFeverIF {
	return &Command{}
}

func init() {
	BuiltinExecutors[base.COMMAND] = NewCommand()
}

func (c *Command) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	length_of_args := codes[pc].(int)

	part_of_stack := VM.PopMultiStack(length_of_args)

	var command []string

	for _, stack := range part_of_stack {
		command = append(command, stack.Val.(string))
	}

	var result []byte
	var err error

	switch len(command) {
	case 1:
		result, err = exec.Command(command[0]).Output()
	default:
		result, err = exec.Command(command[0], command[1:]...).Output()
	}

	if err != nil {
		return pc, err
	}

	VM.PushStack(base.MakeString(string(result)))

	return pc, nil
}
