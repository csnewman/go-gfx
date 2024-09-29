package main

import (
	"log"

	"github.com/csnewman/go-gfx/internal/appkit"
)

func main() {
	log.Println("hello world")

	err := appkit.Run(appkit.Callbacks{
		Init: func() {
			log.Println("init")

			if err := appkit.NewWindow(800, 600); err != nil {
				panic(err)
			}
		},
	})
	if err != nil {
		panic(err)
	}
}
