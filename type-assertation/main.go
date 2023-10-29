package main

import "fmt"

func main() {
	var minhaVar interface{} = "Daniel Paulino"

	println(minhaVar.(string))

	res, ok := minhaVar.(int) // <-- Aqui ele está tentando converter, is called type assertion
	fmt.Printf("O VALOR de RES é %v e o resultado de OK é = %v  ", res, ok)
	res2 := minhaVar.(int) // <-- Here, as it doesn't separate the results, like result of conversion.
	fmt.Printf(" O valor de res 2 é %v", res2)

}
