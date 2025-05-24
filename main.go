package main

import (
	"littlejumbo/genius/managers"

	gomesengine "github.com/mikabrytu/gomes-engine"
)

func main() {
	gomesengine.Init("Genius", 430, 430)
	managers.Game()
	gomesengine.Run()
}
