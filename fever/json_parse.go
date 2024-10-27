package fever

import (
	"disco/base"
	"encoding/json"
)

type JsonParse struct{}

func NewJsonParse() BuiltinFeverIF {
	return &JsonParse{}
}

func init() {
	BuiltinExecutors[base.JSON_PARSE] = NewJsonParse()
}

func (j *JsonParse) Execute(
	codes []any,
	pc int,
	env *base.Environment,
	caller any,
) (int, error) {

	s := VM.PopStack()

	var json_to_s *base.S

	switch s.Val.(string)[0:1] {
	case "[":
		var json_map []map[string]any

		err := json.Unmarshal([]byte(s.Val.(string)), &json_map)
		if err != nil {
			return pc, err
		}

		json_to_s = j.jsonArrayParse(json_map)

	default:
		var json_map map[string]any

		err := json.Unmarshal([]byte(s.Val.(string)), &json_map)
		if err != nil {
			return pc, err
		}

		json_to_s = j.jsonParse(json_map)
	}

	VM.PushStack(json_to_s)

	return pc, nil
}

func (j *JsonParse) jsonArrayParse(json_array []map[string]any) *base.S {
	var json_to_s *base.S

	switch len(json_array) {
	case 0:
		json_to_s = base.NilAtom

	case 1:
		json_to_s = base.Cons(
			j.jsonParse(json_array[0]),
			base.NilAtom,
		)
	default:
		json_to_s = base.Cons(
			j.jsonParse(json_array[0]),
			j.jsonArrayParse(json_array[1:]),
		)
	}

	return json_to_s
}

func (j *JsonParse) jsonParse(m map[string]any) *base.S {
	s := base.NilAtom

	for key, v := range m {
		switch v.(type) {
		case float64:
			s = base.Cons(
				base.Cons(
					base.MakeString(key),
					base.MakeFloat(v.(float64)),
				),
				s,
			)

		case string:
			s = base.Cons(
				base.Cons(
					base.MakeString(key),
					base.MakeString(v.(string)),
				),
				s,
			)

		case []any:
			s = base.Cons(
				base.Cons(
					base.MakeString(key),
					j.valueToVector(v.([]any)),
				),
				s,
			)

		case map[string]any:
			s = base.Cons(
				base.Cons(
					base.MakeString(key),
					j.jsonParse(v.(map[string]any)),
				),
				s,
			)
		}
	}

	return s
}

func (j *JsonParse) valueToVector(value []any) *base.S {
	var vector []*base.S

	for _, v := range value {
		switch v.(type) {
		case float64:
			vector = append(vector, base.MakeFloat(v.(float64)))

		case string:
			vector = append(vector, base.MakeString(v.(string)))

		case []any:
			vector = append(vector, j.valueToVector(v.([]any)))

		case map[string]any:
			vector = append(vector, j.jsonParse(v.(map[string]any)))
		}
	}

	return base.MakeVector(vector)
}
