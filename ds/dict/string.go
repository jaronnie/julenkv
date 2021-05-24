package dict

type Dict struct {
	dict map[string]string
}

func (s *Dict) Set(key, value []byte) error {
	s.dict[string(key)] = string(value)
	return nil
}

func (s *Dict) Get(key []byte) []byte {
	return []byte(s.dict[string(key)])
}

func NewDict() *Dict {
	return &Dict{make(map[string]string)}
}
