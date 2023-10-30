package math

// Soma - When the function is named with the first uppercase letter, it means that is a function can be seen in another files, looks the keyword public from java
// Soma - When the first letter is lowercase means that function can be seen only in your file
func Soma[T float64 | int](num1 T, num2 T) T {
	println("This a can see here, his file named mat", a)
	println("This A can see here, his file named mat as well", A)

	return num1 + num2
}

var A int = 2
var a int = 3

// Carro That rule about the first letter serves for structs as well. So this Carro can be seen in another file
type Carro struct {
	Name string
}

func (c Carro) AndarCarro() {
	println("Tendatndo ver o nome: ", c.Name)
	//println("Tendatndo ver o price: ", c.price)
}
