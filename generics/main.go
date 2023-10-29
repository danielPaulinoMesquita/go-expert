package main

// SomaInt an SomaFloat This can be an overload of methods at Java
func SomaInt(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}

	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}

	return soma
}

// Soma But these two methods above can be transformed in one method named Generics. One sample below
func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

// Number You also can use constraints, probably other types, such as int and float64 extends from Number
// "~" to serve for MyNumber getting compile
type Number interface {
	~int | float64
}

func SomaUsingNumber[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

// MyNumber is a new type of int, but if you pass this type to func that his parameter is Number,
// it probably will get error, so for you use this func, it will make a change on int to ~int of Number constraint,
type MyNumber int

func main() {
	m := map[string]float64{"teste": 12.00, "second": 30.00}
	m2 := map[string]int{"teste": 12, "second": 30}
	m3 := map[string]MyNumber{"testing My": 1, "2": 5}
	println(SomaFloat(m))
	println(SomaInt(m2))
	println(Soma(m2)) // could be passing "m"
	println(SomaUsingNumber(m))
	println(SomaUsingNumber(m3))
}
