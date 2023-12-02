package main

import "github.com/devfullcycle/dan/goexpert/configs"

func main() {
	//...
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
