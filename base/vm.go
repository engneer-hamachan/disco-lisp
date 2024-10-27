package base

import (
	"fmt"
	"os"
	"time"
)

func NewFeverMachine() *FeverMachine {
	var codes []any
	var stack = make([]*S, 0)

	return &FeverMachine{
		stack:       stack,
		Codes:       codes,
		Intaractive: false,
	}
}

type FeverMachine struct {
	stack       []*S
	fstack      []*F
	estack      [][]*S
	Codes       []any
	Intaractive bool
	Time        time.Time
}

func (fm *FeverMachine) EvacutionStack() {
	fm.estack = append(fm.estack, fm.stack)
}

func (fm *FeverMachine) RelocationStack() {
	stack := fm.estack[len(fm.estack)-1]
	fm.estack = fm.estack[:len(fm.estack)-1]

	fm.stack = stack
}

func (fm *FeverMachine) PushStack(s *S) {
	fm.stack = append(fm.stack, s)
}

func (fm *FeverMachine) PeekStack() *S {
	stack := fm.stack[len(fm.stack)-1]

	return stack
}

func (fm *FeverMachine) PopStack() *S {
	stack := fm.stack[len(fm.stack)-1]
	fm.stack = fm.stack[:len(fm.stack)-1]

	return stack
}

func (fm *FeverMachine) PopMultiStack(lengs_of_args int) []*S {
	part_of_stack := fm.stack[len(fm.stack)-lengs_of_args:]
	fm.stack = fm.stack[:len(fm.stack)-lengs_of_args]

	return part_of_stack
}

func (fm *FeverMachine) Fatal(caller string, err error, pc int) {
	info, ok := InformationWhenParsing[caller]
	if !ok {
		info = InformationWhenCompile[caller][pc]
	}

	fmt.Printf("%v::", info.FileName)
	fmt.Printf("%v::", info.Row)
	fmt.Printf("%v\n", err)

	if !fm.Intaractive {
		os.Exit(1)
	}
}

func (fm *FeverMachine) Dump() {
	fmt.Println("=======DUMP START========")

	for _, byte_code := range fm.Codes {
		dumpByteCode(byte_code)
	}

	fmt.Println("=======DUMP END========")
}
