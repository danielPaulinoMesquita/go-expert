package main

import (
	"fmt"
	"github.com/google/uuid"
	"golandProjects/awesomeProject/packaging/math"
)

func main() {
	// to import a package from local and not in the external repository, you need to use -replace
	// like: go mod edit -replace github.com/danielPaulinoMesquita/go-expert/packaging/math=../
	// this is very useful to test modules before to release
	println(uuid.New().String())
	m := math.Math{A: 2, B: 3}
	fmt.Println(m.Add())

}
