package base

var Globals map[string]*S = make(map[string]*S)

type Environment struct {
	Stack []map[string]*S
	Peek  map[string]*S
}

func NewEnvironment() Environment {
	return Environment{Stack: make([]map[string]*S, 0)}
}

func (e *Environment) size() int {
	return len(e.Stack)
}

func (e *Environment) tailIdx() int {
	return e.size() - 1
}

func (e *Environment) PushStack() {
	var t map[string]*S = make(map[string]*S)

	e.Stack = append(e.Stack, t)
	e.Peek = t
}

func (e *Environment) PopStack() {
	e.Stack = e.Stack[:e.tailIdx()]

	if len(e.Stack) > 0 {
		e.Peek = e.Stack[e.tailIdx()]
	}
}

func (e *Environment) GetSymbolValue(symbol *S) (*S, bool) {
	if e.size() != 0 {
		symbol_name := symbol.Val.(string)

		for tail_idx := e.tailIdx(); tail_idx >= 0; tail_idx-- {
			v, ok := e.Stack[tail_idx][symbol_name]
			if ok {
				return v, ok
			}
		}
	}

	val, ok := Globals[symbol.Val.(string)]
	return val, ok
}

func (e *Environment) SetSymbolValue(symbol *S, value *S) {
	if e.size() == 0 {
		Globals[symbol.Val.(string)] = value
		return
	}

	symbol_name := symbol.Val.(string)

	for tail_idx := e.tailIdx(); tail_idx >= 0; tail_idx-- {
		t := e.Stack[tail_idx]

		if _, ok := t[symbol_name]; ok {
			t[symbol_name] = value
			return
		}
	}

	e.Peek[symbol_name] = value
}

func (e *Environment) SetGlobalSymbolValue(symbol *S, value *S) {
	Globals[symbol.Val.(string)] = value
}

func (e *Environment) SetSymbolValueForPeek(symbol *S, value *S) {
	e.Peek[symbol.Val.(string)] = value
}

func (e *Environment) SetMultiSymbolValueForPeek(dargs []*S, args []*S) {
	for idx, arg := range args {
		e.Peek[dargs[idx].Val.(string)] = arg
	}

	if len(dargs) > len(args) {
		for _, darg := range dargs[len(args):] {
			e.Peek[darg.Val.(string)] = NilAtom
		}
	}
}
