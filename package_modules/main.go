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

	println("Valor é: ", s)
}
