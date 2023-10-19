package main

import (
	"errors"
	"fmt"
)

func main() {
	result := sum(4, 5)
	fmt.Println("Resultado: ", result)

	fmt.Println(sum2(4, 6))
	fmt.Println(sum2(4, 3)) // showed up the error nil

	// catch the value and error
	resultNew, err := sum2(8, 3)
	if err != nil {
		fmt.Println("Capturou o erro:", err)
	} else {
		fmt.Println("Sem error")
		fmt.Println(resultNew)
	}
}

// first way to create a function (bad alternative)
func sum(int2 int, int1 int) int {
	return int2 + int1
}

// second way to create a function (better alternative)
func sum2(num1, num2 int) (int, error) {
	if num1+num2 > 8 {
		return 0, errors.New("erro, soma é maior que 8")
	}
	// nil looks like null
	return num1 + num2, nil
}
