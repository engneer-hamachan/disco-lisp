package printer

import (
	"disco/base"
	"disco/predicater"
	"fmt"
	"strings"
)

func Print(s *base.S, is_princ bool) {
	switch s.Type {
	case base.INT:
		fmt.Print(s.Val.(int64))

	case base.FLOAT:
		fmt.Print(s.Val.(float64))

	case base.STRING:
		if is_princ {
			fmt.Printf("%v", s.Val.(string))
			break
		}

		fmt.Printf("\"%v\"", s.Val.(string))

	case base.SYMBOL, base.TRUE, base.NIL:
		fmt.Print(strings.ToUpper(s.Val.(string)))

	case base.LIST:
		fmt.Print("(")
		printList(s, is_princ)

	case base.VECTOR:
		fmt.Print("#VECTOR(")
		printVector(s, is_princ)

	case base.HASH:
		fmt.Print("#TABLE")
	}
}

func printList(s *base.S, is_princ bool) {
	if predicater.Nilp(s) {
		fmt.Print(")")
		return
	}

	if predicater.Dotlistp(s) {
		Print(s.GetCar(), is_princ)
		fmt.Print(" . ")

		Print(s.GetCdr(), is_princ)
		fmt.Print(")")

		return
	}

	Print(s.GetCar(), is_princ)

	if !predicater.Nilp(s.GetCdr()) {
		fmt.Print(" ")
	}

	printList(s.GetCdr(), is_princ)
}

func printVector(s *base.S, is_princ bool) {
	vector := s.Val.([]*base.S)

	for idx, vec := range vector {
		Print(vec, is_princ)

		if idx != len(vector)-1 {
			fmt.Print(" ")
		}
	}

	fmt.Print(")")
}
