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

			if _, err := appkit.NewWindow(800, 600); err != nil {
				panic(err)
			}
		},
		CloseRequested: func(w appkit.Window) {
			log.Println("close requested", w)

			w.Close()
		},
		Closed: func(w appkit.Window) {
			log.Println("closed", w)

			appkit.Stop()
		},
	})
	if err != nil {
		panic(err)
	}

	log.Println("main returned")
}
