package main

const a = "Hellow World"

// ID Creating a type of variable
type ID int

var (
	b bool = true
	c int
	d string
	e float64 = 1.2
	f ID      = 21
)

func main() {

	// x := "shorthand" <- if variable local is declared and not used,
	// the compilation go will warn (x declared and not used)

	//println(a)
	//println(b)
	//println(c)
	//println(d)
	//println(x)
	println(f)
}
