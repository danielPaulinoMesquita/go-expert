package main

import "fmt"

func main() {

	var meuArray [3]int
	meuArray[0] = 11
	meuArray[1] = 222
	meuArray[2] = 312

	fmt.Println(meuArray[0])
	fmt.Println(len(meuArray))             //<-size or length
	fmt.Println(meuArray[len(meuArray)-1]) //<- getting value using len(size)
	// (this is a good choice to discover and not fail when we try to get a value) minus one position

	for i, v := range meuArray {
		fmt.Printf("The Value of index is %d and the value is %d\n", i, v)
		//fmt.Printf("The Value of index is %d and the value is %s case it was string \n", i, v)
	}

}
