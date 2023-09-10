package source

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PyMarcus/go_gaming/events"
	"github.com/PyMarcus/go_gaming/objects"
	"github.com/PyMarcus/go_gaming/settings"
	"github.com/veandco/go-sdl2/sdl"
)

var playerManager = make(map[string]*objects.Player)

// PlayGame starts the game and settings of screen
func PlayGame(playerName string) {
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

	keepRunningGameWindow(playerName, renderer)

	os.Exit(0)
}

// keepGameWindow keeps the window on load
func keepRunningGameWindow(playerName string, renderer *sdl.Renderer) {
	for {
		frameStartTime := time.Now()

		if !closeWindow() {
			return
		}

		renderer.Clear()
		delta := int(time.Since(frameStartTime).Seconds()) * settings.FPS

		currentPlayer, found := playerManager[playerName]

		if !found {
			playerManager[playerName] = objects.NewPlayerOnline(settings.PLAYER_IMAGE_PATH, 0, 0, renderer)
			currentPlayer = playerManager[playerName]
		}
		name, px, py := mqtt(playerName, currentPlayer.Posx, currentPlayer.Posy)

		if name != playerName {
			newPlayer, found := playerManager[name]
			if !found {
				newPlayer = objects.NewPlayerOnline(settings.PLAYER_IMAGE_PATH, px, py, renderer)
				playerManager[name] = newPlayer
			}

			/*
			playerName = name
			delete(playerManager, playerName)
			currentPlayer = newPlayer
			*/
		}
		
		fmt.Println("A renderizar ", len(playerManager), playerManager)

		for _, player := range playerManager {
			player.Draw(renderer, delta)
		}

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

func mqtt(playerName string, posx, posy float64) (string, float64, float64){
	response := events.PubAndRecv(playerName, posx, posy)
	received := strings.Split(response, ":")
	name := received[0]
	if len(received) > 2{
		fmt.Println("ok")
		posxx := received[1]
		posyy := received[2]
		x, _ := strconv.ParseFloat(posxx, 64)
		y, _ := strconv.ParseFloat(posyy, 64)
		return name, x, y
	}
	return "ausente", 0.0, 0.0
}
