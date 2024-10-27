package base

type S struct {
	Type     int
	Val      any
	car      *S
	cdr      *S
	IsQuoted bool
}

func NewS(
	types int,
	val any,
	car *S,
	cdr *S,
) *S {

	return &S{
		Type:     types,
		Val:      val,
		car:      car,
		cdr:      cdr,
		IsQuoted: false,
	}
}

func MakeLoad() *S {
	return NewS(LOAD, nil, NilAtom, NilAtom)
}

func MakeFunc(val any, args *S) *S {
	return NewS(FUNC, val, args, nil)
}

func MakeExecFunc(val any) *S {
	return NewS(EXEC_FUNC, val, nil, nil)
}

func MakeMacro(val any, args *S, body *S) *S {
	return NewS(MACRO, val, args, body)
}

var int64Cache = make(map[int64]*S)

func MakeInt(val int64) *S {
	cache_s, ok := int64Cache[val]
	if ok {
		switch cache_s.Val.(type) {
		case int64:
			if cache_s.Val.(int64) == val {
				cache_s.IsQuoted = false
				return cache_s
			}

		default:
			break
		}
	}

	s := NewS(INT, val, NilAtom, NilAtom)
	int64Cache[val] = s

	return s
}

func MakeFloat(val any) *S {
	return NewS(FLOAT, val, NilAtom, NilAtom)
}

func MakeString(val any) *S {
	return NewS(STRING, val, NilAtom, NilAtom)
}

func MakeVector(val any) *S {
	return NewS(VECTOR, val, NilAtom, NilAtom)
}

func MakeHash(val any) *S {
	return NewS(HASH, val, NilAtom, NilAtom)
}

func MakeFp(val any) *S {
	return NewS(FP, val, NilAtom, NilAtom)
}

func MakeReq(val any) *S {
	return NewS(REQUEST, val, NilAtom, NilAtom)
}

func MakeTrue() *S {
	return NewS(TRUE, "t", NilAtom, NilAtom)
}

func MakeNil() *S {
	return NewS(
		NIL,
		"nil",
		NewS(NIL, "nil", nil, nil),
		NewS(NIL, "nil", nil, nil),
	)
}

func MakeDummyS(base_type int, val any) *S {
	return NewS(base_type, val, NilAtom, NilAtom)
}

var symCache = make(map[string]*S)

func MakeSym(val string) *S {
	cache_s, ok := symCache[val]
	if ok && cache_s.Val.(string) == val {
		cache_s.IsQuoted = false
		return cache_s
	}

	sym := NewS(SYMBOL, val, NilAtom, NilAtom)

	return sym
}

func Cons(car *S, cdr *S) *S {
	return NewS(LIST, "list-data", car, cdr)
}

func Append(car *S, cdr *S) *S {
	if car.Type == NIL {
		return cdr
	}

	return Cons(
		car.GetCar(),
		Append(car.GetCdr(), cdr),
	)
}

func (s *S) GetCar() *S {
	return s.car
}

func (s *S) GetCdr() *S {
	return s.cdr
}

func (s *S) GetCaar() *S {
	return s.car.car
}

func (s *S) GetCadr() *S {
	return s.cdr.car
}

func (s *S) GetCdar() *S {
	return s.car.cdr
}

func (s *S) GetCddr() *S {
	return s.cdr.cdr
}

func (s *S) GetCadar() *S {
	return s.car.cdr.car
}

func (s *S) GetCaddr() *S {
	return s.cdr.cdr.car
}

func (s *S) GetCadddr() *S {
	return s.cdr.cdr.cdr.car
}

func (s *S) GetCdadr() *S {
	return s.cdr.car.cdr
}

func (s *S) GetCadadr() *S {
	return s.cdr.car.cdr.car
}

func (s *S) GetCaadr() *S {
	return s.cdr.car.car
}

func (s *S) SetCar(newS *S) {
	s.car = newS
}

func (s *S) SetCdr(newS *S) {
	s.cdr = newS
}

var NilAtom = MakeNil()
var TrueAtom = MakeTrue()
var Load = MakeLoad()
