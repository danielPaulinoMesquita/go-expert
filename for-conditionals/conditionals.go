package main

func If() {
	a := 1
	b := 3

	// conditionals
	// ==
	// >=
	// <=
	// <
	// you can use &&(AND) and ||(OR)

	if a > b {
		println("maior é o: ", a)
	} else {
		println("Menor é o b:", b)
	}

	c := 1
	//d := 2
	//e := 3

	switch c {
	case 1:
		println("a")
	case 2:
		println("b")
	default:
		println("Qualquer outra coisa")
	}

}
