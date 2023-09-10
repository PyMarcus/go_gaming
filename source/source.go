package source

import (
	"fmt"
	"github.com/PyMarcus/go_gaming/events"
	"github.com/PyMarcus/go_gaming/objects"
	"github.com/PyMarcus/go_gaming/settings"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/ttf"
	"os"
	"strconv"
	"strings"
	"time"
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
	ttf.Init()

	font, err := ttf.OpenFont("/usr/share/fonts/TTF/AkaashNormal.ttf", 24)
	ttf.Init()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao carregar a fonte: %v\n", err)
		os.Exit(1)
	}
	
	defer font.Close()

	currentPlayer, found := playerManager[playerName]

	if !found {
		playerManager[playerName] = objects.NewPlayerOnline(playerName, settings.PLAYER_IMAGE_PATH, 0, 0, renderer)
		currentPlayer = playerManager[playerName]
	}
	
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar o renderer: %v\n", err)
		os.Exit(1)
	}
	defer renderer.Destroy()

	

	for {
		frameStartTime := time.Now()
		delta := int(time.Since(frameStartTime).Seconds()) * settings.FPS

		if !closeWindow() {
			return
		}

		renderer.Clear()
		
		backgroundImage, err := img.LoadTexture(renderer, settings.BACKGROUND)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao carregar a imagem: %v\n", err)
			os.Exit(1)
		}
		defer backgroundImage.Destroy()
		renderer.Copy(backgroundImage, nil, nil)


		name, px, py := mqtt(playerName, currentPlayer.Posx, currentPlayer.Posy)

		if name != playerName && name != "ausente" && playerName != "ausente" {
			newPlayer, found := playerManager[name]
			if !found {
				newPlayer = objects.NewPlayerOnline(name, settings.PLAYER_IMAGE_PATH, px, py, renderer)
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
			
			textSurface, err := font.RenderUTF8Blended(player.Name, sdl.Color{R: 255, G: 0, B: 0, A: 0})
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao renderizar o texto: %v\n", err)
				os.Exit(1)
			}
			defer textSurface.Free()

			textTexture, err := renderer.CreateTextureFromSurface(textSurface)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro ao criar a textura do texto: %v\n", err)
				os.Exit(1)
			}
			defer textTexture.Destroy()

			textRect := &sdl.Rect{X: int32(player.Posx), Y: int32(player.Posy - 20), W: textSurface.W, H: textSurface.H}
			renderer.Copy(textTexture, nil, textRect)
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

func mqtt(playerName string, posx, posy float64) (string, float64, float64) {
	response := events.PubAndRecv(playerName, posx, posy)
	received := strings.Split(response, ":")
	name := received[0]
	if len(received) > 2 {
		posxx := received[1]
		posyy := received[2]
		x, _ := strconv.ParseFloat(posxx, 64)
		y, _ := strconv.ParseFloat(posyy, 64)
		return name, x, y
	}
	return "ausente", 0.0, 0.0
}
