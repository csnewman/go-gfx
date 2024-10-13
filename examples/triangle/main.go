package main

import (
	"github.com/csnewman/go-gfx/gfx"
	"log"
	"time"
)

var lastPrint time.Time
var count int

type Example struct {
	app    *gfx.Application
	window *gfx.Window
}

func (e *Example) init(app *gfx.Application) error {
	e.app = app

	log.Println("Creating main window")

	window, err := e.app.NewWindow(gfx.WindowConfig{
		Title:    "Triangle",
		Width:    800,
		Height:   600,
		OnClosed: e.closed,
		OnRender: e.render,
	})
	if err != nil {
		return err
	}

	e.window = window

	return nil
}

func (e *Example) Run() error {
	return gfx.Run(gfx.ApplicationConfig{
		Init: e.init,
	})
}

func (e *Example) closed() {
	log.Println("Main window closed")

	e.app.Exit()
}

func (e *Example) render() {
	count++

	if time.Since(lastPrint) > time.Second {
		lastPrint = time.Now()

		log.Println("FPS", count)
		count = 0
	}
}

func main() {
	log.Println("Triangle Example")

	ex := &Example{}

	if err := ex.Run(); err != nil {
		panic(err)
	}

	log.Println("App exited")
}
