package main

import "fmt"

func main() {

	anyArray := []int{11, 22, 33, 44, 55}

	// cap means capacity, in this case is anyArray
	// len means size of array

	fmt.Printf("Testing || Tamanho (len)=%d || Capacidade(cap)=%d || Valores(values)=%v \n", len(anyArray), cap(anyArray), anyArray)
	fmt.Printf("\nTesting SLICE without modifications \n|| Tamanho (len)=%d || Capacidade(cap)=%d || Valores(values)=%v", len(anyArray[0:]), cap(anyArray[0:]), anyArray[:0])
	fmt.Println("\n-----------------------------------------------------------")
	fmt.Printf("Testing SLICE with modification \n|| Tamanho (len)=%d || Capacidade(cap)=%d || Valores(values)=%v", len(anyArray[0:2]), cap(anyArray[2:]), anyArray[2:])

	// add new value, or increase capacity
	anyArray = append(anyArray, 66)
	//Now, it duplicates capacity, looks like at the array from java.
	//This duplicate happens because it has a size fixed with initial amount values (or the size that is passed it)
	//and when is put a new value into array to be capable of receive this new value, the array duplicates,
	//his rule, don't ask me about that
	fmt.Println("\n-----------------------------------------------------------")

	fmt.Printf("Testing adden new value || Tamanho (len)=%d || Capacidade(cap)=%d || Valores(values)=%v \n", len(anyArray), cap(anyArray), anyArray)

}
