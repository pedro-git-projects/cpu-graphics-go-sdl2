package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	window    *sdl.Window
	renderer  *sdl.Renderer
	isRunning bool = false
)

func initializeWindow() bool {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Printf("Error initilizing SDL %s", err)
		return false
	}

	window, err = sdl.CreateWindow("", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 800, 600, sdl.WINDOW_BORDERLESS)
	if err != nil {
		fmt.Printf("Error initilizing window %s", err)
		return false
	}

	renderer, err = sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		fmt.Printf("Error initilizing renderer %s", err)
		return false
	}
	return true
}

func processInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			isRunning = false
		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				isRunning = false
			}
		}
	}
}

func render() {
	renderer.SetDrawColor(255, 25, 25, 255)
	renderer.Clear()
	renderer.Present()
}

func main() {
	isRunning = initializeWindow()

	for isRunning {
		processInput()
		render()
	}
}
