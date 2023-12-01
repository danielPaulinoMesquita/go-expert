package main

import (
	"fmt"
	"github.com/google/uuid"
	"golandProjects/awesomeProject/packaging/math"
)

func main() {
	m := math.Math{A: 2, B: 3}

	fmt.Println(m.Add())
	fmt.Println(" Hellow, World! ")
	fmt.Println(uuid.New().String())
}
