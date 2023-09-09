package objects

import (
	"fmt"
	
	"github.com/PyMarcus/go_gaming/settings"
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	Texture *sdl.Texture
}

var count int8 = 0

func NewPlayer(imagePath string, renderer *sdl.Renderer) (Player) {
	texture := createPlayerTexture(imagePath, renderer)
	return Player{Texture: texture}
}

// createPlayerTexture textures are most performatics
func createPlayerTexture(imagePath string, renderer *sdl.Renderer) *sdl.Texture {
	img, err := sdl.LoadBMP(imagePath)
	defer img.Free()

	if err != nil {
		fmt.Println("Falha ao carregar imagem bmp do player")
		panic(err)
	}

	playerTexture, err := renderer.CreateTextureFromSurface(img)

	return playerTexture
}

func (p *Player) Draw(renderer *sdl.Renderer){
	if count == 0{
		renderer.Copy(p.Texture, &sdl.Rect{
			X: settings.PLAYER_START_POSITION_X,
			Y: settings.PLAYER_START_POSITION_Y,
			W: settings.PLAYER_SIZE_WIDTH,
			H: settings.PLAYER_SIZE_HEIGTH},
			&sdl.Rect{
				X: settings.PLAYER_START_POSITION_X,
				Y: settings.PLAYER_START_POSITION_Y,
				W: settings.PLAYER_SIZE_WIDTH,
				H: settings.PLAYER_SIZE_HEIGTH},
		)
	}else{
		count = 1
	}
}