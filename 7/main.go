package main

import "fmt"

func main() {
	mapNames := map[string]int{"1- um": 1, "2- dois": 2, "3- tres": 3} // first way to create map, initializing map
	fmt.Println(mapNames)
	delete(mapNames, "2- dois")
	fmt.Println(mapNames)
	fmt.Println(mapNames["3- tres"])
	mapNames["5- cinco"] = 5

	var mapNamesOld = make(map[int]string) // second way to create a map
	mapNamesOld[1] = "teste one"
	mapNamesOld[2] = "teste dois"

	var mapNamesNew = map[int]string{} // third way to create, initializing map with empty
	mapNamesNew[4] = "teste quatro"
	fmt.Println(mapNamesOld)

	for chave, valor := range mapNames {
		fmt.Printf("Verificando chaves:(%s) valores=%d\n", chave, valor)
	}

	// here, we are using blank identifier (it's used to avoid some value)
	for _, valor := range mapNames {
		fmt.Printf("Verificando valores=%d\n", valor)
	}
}
