package fever

import (
	"disco/base"
	"fmt"
	"net/http"
)

type DefHandler struct{}

var queries map[string]*base.S
var writer http.ResponseWriter

func NewDefHandler() BuiltinFeverIF {
	return &DefHandler{}
}

func init() {
	BuiltinExecutors[base.DEFHANDLER] = NewDefHandler()
}

func (d *DefHandler) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	f := VM.PopStack()
	s := VM.PopStack()

	handler := func(w http.ResponseWriter, r *http.Request) {
		writer = w

		d.setQueries(r)

		Fever(f.Val.(*base.F).Body, env, s.Val.(string))

		content := VM.PopStack()

		fmt.Fprint(w, content.Val.(string))
	}

	http.HandleFunc(s.Val.(string), handler)

	VM.PushStack(base.TrueAtom)

	return pc, nil
}

func (d *DefHandler) setQueries(
	r *http.Request,
) {
	queries = make(map[string]*base.S)

	m := r.URL.Query()

	for k, v := range m {
		queries[k] = base.MakeString(v[0])
	}
}
