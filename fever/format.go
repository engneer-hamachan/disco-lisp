package fever

import (
	"disco/base"
	"strconv"
	"strings"
)

type Format struct{}

func NewFormat() BuiltinFeverIF {
	return &Format{}
}

func init() {
	BuiltinExecutors[base.FORMAT] = NewFormat()
}

func (f *Format) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	pc += 1

	length_of_args := codes[pc].(int)
	part_of_stack := VM.PopMultiStack(length_of_args)
	text := part_of_stack[0]

	var err error

	text, err = f.format(text, part_of_stack)
	if err != nil {
		return pc, err
	}

	VM.PushStack(text)

	return pc, nil
}

func (f *Format) format(
	text *base.S,
	part_of_stack []*base.S,
) (*base.S, error) {

	text = base.MakeString(
		strings.ReplaceAll(
			text.Val.(string),
			"~%",
			"\n",
		),
	)

	if len(part_of_stack) < 2 {
		return text, nil
	}

	args := part_of_stack[1:]

	for _, a := range args {
		switch a.Val.(type) {
		case string:
			text = base.MakeString(
				strings.Replace(
					text.Val.(string),
					"~a",
					a.Val.(string),
					1,
				),
			)

		case int64:
			text = base.MakeString(
				strings.Replace(
					text.Val.(string),
					"~d",
					strconv.FormatInt(a.Val.(int64), 10),
					1,
				),
			)

		case float64:
			text = base.MakeString(
				strings.Replace(
					text.Val.(string),
					"~d",
					strconv.FormatFloat(
						a.Val.(float64),
						'f',
						-1,
						64,
					),
					1,
				),
			)
		}
	}

	return text, nil
}
