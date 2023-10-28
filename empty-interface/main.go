package main

import "fmt"

/*
*
looks like generics
*/
func main() {

	// This is like a Generics from java,
	// generic type that be uses the same actions that others
	var x interface{} = 10
	var y interface{} = "Hello World!"

	showType(x)
	showType(y)

}

func showType(t interface{}) {
	// ''t'' is the object that can be anything,
	// but we want to make the same actions
	fmt.Printf("O tipo é %T e o valor é %v\n", t, t)
}
