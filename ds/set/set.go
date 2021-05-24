package set

type Record map[string]map[string] bool

type Set struct {
	record Record
}

func New() *Set {
	return &Set{make(Record)}
}

