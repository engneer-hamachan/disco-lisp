package compiler

import "disco/base"

func sLength(s *base.S) int {
	if s.GetCar() == nil || (s.GetCar().Type == base.NIL && s.GetCdar() == nil) {
		return 0
	}

	ct := 1
	cdr := s.GetCdr()

	for {
		if cdr.Type == base.NIL {
			return ct
		}

		ct += 1

		cdr = cdr.GetCdr()
	}
}

func sLengthWithoutOptional(s *base.S) int {
	if s.GetCar() == nil ||
		(s.GetCar().Type == base.NIL && s.GetCdar() == nil) ||
		s.GetCar().Val.(string) == "&optional" {

		return 0
	}

	ct := 1
	cdr := s.GetCdr()

	for {
		if cdr.Type == base.NIL || cdr.GetCar().Val.(string) == "&optional" {
			return ct
		}

		ct += 1

		cdr = cdr.GetCdr()
	}
}

func sLengthWithOptional(s *base.S) int {
	if s.GetCar() == nil || (s.GetCar().Type == base.NIL && s.GetCdar() == nil) {
		return 0
	}

	ct := 1
	cdr := s.GetCdr()

	for {
		if cdr.Type == base.NIL {
			return ct
		}

		if cdr.GetCar().Val.(string) != "&optional" {
			ct += 1
		}

		cdr = cdr.GetCdr()
	}
}
