package math

// Math the variables uppercase is necessary for viewing in another file go
type Math struct {
	A int
	B int
}

func (m Math) Add() int {
	return m.A + m.B
}
