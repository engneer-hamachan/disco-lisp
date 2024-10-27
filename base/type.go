package base

var tbl map[string]Symbol = make(map[string]Symbol)

type Symbol struct {
	name string
}

func newSymbol(name string) Symbol {
	symbol := Symbol{
		name: name,
	}

	tbl[name] = symbol

	return symbol
}

func (s *Symbol) GetName() string {
	return s.name
}

func Intern(name string) Symbol {
	val, ok := tbl[name]
	if ok {
		return val
	}

	return newSymbol(name)
}

const (
	EOS           = -1
	NIL           = 0
	INT           = 257
	SYMBOL        = 258
	STRING        = 259
	TRUE          = 260
	LIST          = 261
	BUILTIN       = 262
	LOAD          = 263
	FUNC          = 264
	FLOAT         = 265
	MACRO         = 266
	ANY           = 267
	VECTOR        = 268
	FP            = 269
	REQUEST       = 270
	QUOTED_SYMBOL = 271
	HASH          = 272
	OPTIONAL      = 273
)

func TypeToString(base_type int) string {
	switch base_type {
	case NIL:
		return "nil"
	case INT:
		return "int"
	case SYMBOL:
		return "symbol"
	case STRING:
		return "string"
	case TRUE:
		return "true"
	case LIST:
		return "list"
	case BUILTIN:
		return "builtin"
	case LOAD:
		return "load"
	case FUNC:
		return "func"
	case EXEC_FUNC:
		return "exec-func"
	case FLOAT:
		return "float"
	case MACRO:
		return "macro"
	case ANY:
		return "any"
	case FP:
		return "file pointer"
	case REQUEST:
		return "request"
	case VECTOR:
		return "vector"
	case QUOTED_SYMBOL:
		return "quoted symbol"
	case HASH:
		return "hash table"
	case OPTIONAL:
		return "optional"
	default:
		panic("type convert error")
	}
}

func OptionalTypeToString(optional_type []int) string {
	str := "optional<"

	for _, types := range optional_type {
		str += TypeToString(types)
		str += " "
	}

	str = str[:len(str)-1]
	str += ">"

	return str
}

func IsMatchOptionalType(optional_type []int, target_type int) bool {
	if len(optional_type) == 0 {
		return true
	}

	for _, types := range optional_type {
		if types == target_type {
			return true
		}
	}

	return false
}

func IsMatchOptionalTypeForNumber(optional_type []int) bool {
	if len(optional_type) == 0 {
		return true
	}

	for _, types := range optional_type {
		if types == INT || types == FLOAT {
			return true
		}
	}

	return false
}
