package main

import (
	"fmt"
	"github.com/devfullcycle/fcutils/pkg/events"
)

func main() {
	// it can see more on the "Módulos privados" of fullcycle course, but at this moment I don't see much importance on it
	// but has something related to GOPRIVATE
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
