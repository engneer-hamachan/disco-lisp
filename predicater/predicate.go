package predicater

import (
	"disco/base"
)

func Dotlistp(s *base.S) bool {
	return s.GetCdr().Type != base.LIST
}

func Symbolp(s *base.S) bool {
	return s.Type == base.SYMBOL
}

func Intp(s *base.S) bool {
	return s.Type == base.INT
}

func Floatp(s *base.S) bool {
	return s.Type == base.FLOAT
}

func Stringp(s *base.S) bool {
	return s.Type == base.STRING
}

func Truep(s *base.S) bool {
	return s.Type == base.TRUE
}

func Nilp(s *base.S) bool {
	return s.Type == base.NIL
}

func Atomp(s *base.S) bool {
	return Intp(s) ||
		Floatp(s) ||
		Symbolp(s) ||
		Stringp(s) ||
		Truep(s) ||
		Nilp(s)
}

func Listp(s *base.S) bool {
	return s.Type == base.LIST
}

func Builtinp(s *base.S) bool {
	return s.Type == base.BUILTIN
}

func Funcp(s *base.S) bool {
	return s.Type == base.FUNC
}

func Macrop(s *base.S) bool {
	return s.Type == base.MACRO
}

func HasName(s *base.S, substr string) bool {
	return s.Val.(string) == substr
}
