package main

import "fmt"

func main() {
	evento := []string{"teste", "teste2", "teste3", "teste4"}
	evento2 := append(evento[:0], evento[1:]...) // it will remove the first element and copy all other elements to first positions
	evento = evento[:2]                          // it will show us only teste3 and teste4

	fmt.Println(evento)
	fmt.Println(evento2)

}
