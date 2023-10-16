package main

const a = "Hellow World"

//var b bool
//var c int
//var d string

// variables global
var (
	b bool = true
	c int
	d string
)

func main() {

	var a string = "test variable local " // variable local

	// b = "owerwritevalue" <- this can't happen, because it is strongest typed.

	e := "shorthand"
	//e := false once time that you created and used shorthand, you can't create again

	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
}
