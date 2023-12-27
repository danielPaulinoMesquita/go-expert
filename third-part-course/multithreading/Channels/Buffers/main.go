package main

func main() {
	ch := make(chan string, 2)
	// buffer is the number, with this, you can add and read more channels
	ch <- "Hello"
	ch <- "World"
	//ch <- "3" <- it will throw deadlock because there only 2 buffers

	println(<-ch)
	println(<-ch)
}
