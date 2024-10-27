package base

type Info struct {
	FileName string
	Row      int
}

var InformationWhenCompile = make(map[string]map[int]Info)
var InformationWhenParsing = make(map[any]Info)

func NewInfo(file_name string, row int, code any) Info {
	info_when_parsed, ok := InformationWhenParsing[code]
	if ok {
		return info_when_parsed
	}

	return Info{
		FileName: file_name,
		Row:      row,
	}
}
