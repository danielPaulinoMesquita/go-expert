package main

/**
go mod init github.com/danielPaulinoMesquita/go-expert/package_modules
now, after we run the command line go mod init
we can import others packages in main file
*/
import "github.com/danielPaulinoMesquita/go-expert/package_modules/math"

func main() {
	println("testando main")

	s := math.Soma(2, 6)
	//  s1 := math.soma(2, 6) <--  this isn't possible, because the first word is lowercase, so it is only possible to see in his file, where it was created or declared

	println("Valor é: ", s)

	carro := math.Carro{Name: "Fiate"} // carro.price = 20.00 <- isn't possible

	carro.AndarCarro()
	println("Instance of Carro: ", carro.Name)
	// println("This 'a' cannot see here, unfortunately", math.a)
	println("This A can see here as well", math.A)
}
