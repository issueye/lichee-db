package core

type Str []byte

func (s Str) Get(key string) string {
	return string(s)
}

func (s Str) Put(key string, value string) {
	s = []byte(value)
}
