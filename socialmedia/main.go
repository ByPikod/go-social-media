package main

import (
	"socialmedia/core"
	"socialmedia/router"
)

func main() {
	core.InitDatabase()
	router.Listen()
}
