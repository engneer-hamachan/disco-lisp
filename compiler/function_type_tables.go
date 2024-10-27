package compiler

type MapKeys struct {
	key1 string
	key2 string
}

var FunctionReturnTypes = make(map[string]int)
var FunctionArgumentTypes = make(map[MapKeys]int)
var FunctionOptionalTypes = make(map[string][]int)
var FunctionArgumentOptionalTypes = make(map[MapKeys][]int)
