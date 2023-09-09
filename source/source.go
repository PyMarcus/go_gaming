package source

import (
	"fmt"
	"os"

	"github.com/PyMarcus/go_gaming/objects"
	"github.com/PyMarcus/go_gaming/settings"
	"github.com/veandco/go-sdl2/sdl"
)

// PlayGame starts the game and settings of screen
func PlayGame() {
	// screen settings
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao inicializar SDL: %v\n", err)
		os.Exit(1)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		settings.WINDOW_TITLE,
		settings.INITIAL_SCREEN_POSITION_X,
		settings.INITIAL_SCREEN_POSITION_Y,
		settings.SCREEN_SIZE_X,
		settings.SCREEN_SIZE_Y,
		sdl.WINDOW_SHOWN)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar janela: %v\n", err)
		os.Exit(1)
	}

	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar renderer: %v\n", err)
		os.Exit(1)
	}
	defer renderer.Destroy()

	// game

	keepRunningGameWindow(renderer)

	os.Exit(0)
}

// keepGameWindow keeps the window on load
func keepRunningGameWindow(renderer *sdl.Renderer) {
	for {

		if !closeWindow() {
			return
		}

		renderer.Clear()
		player := objects.NewPlayer(settings.PLAYER_IMAGE_PATH, renderer)
		player.Draw(renderer)

		renderer.Present()
	}
}

// closeWindow closes the main window
func closeWindow() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return false
		}
	}
	return true
}
