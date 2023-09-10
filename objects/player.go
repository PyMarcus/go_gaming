package objects

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/PyMarcus/go_gaming/settings"
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	Texture *sdl.Texture
	Posx    float64 
	Posy    float64
}

var count int8 = 0

func NewPlayer(imagePath string, renderer *sdl.Renderer) (Player) {
	texture := createPlayerTexture(imagePath, renderer)
	return Player{Texture: texture}
}

func NewPlayerOnline(imagePath string, posx, posy float64, renderer *sdl.Renderer) (*Player) {
	fmt.Println("Criando jogador com as seguintes posicoes ", posx, posy)
	texture := createPlayerTexture(imagePath, renderer)
	return &Player{Texture: texture, Posx: posx, Posy: posy}
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

func (p *Player) Draw(renderer *sdl.Renderer, fps int){
	rand.Seed(time.Now().UnixNano())
	
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
		count = 1
		renderer.Clear()
	}
	if count == 1{
		p.update(renderer, fps)
	}
	
	
	
}

func (p *Player) update(renderer *sdl.Renderer, fps int){
	p.Posx = float64(rand.Intn(500)) - settings.PLAYER_SPEED * float64(fps)
	p.Posy = float64(rand.Intn(500)) - settings.PLAYER_SPEED * float64(fps)
	time.Sleep(time.Second * 2)
	renderer.Copy(p.Texture, &sdl.Rect{
		X: settings.PLAYER_START_POSITION_X,
		Y: settings.PLAYER_START_POSITION_Y,
		W: settings.PLAYER_SIZE_WIDTH,
		H: settings.PLAYER_SIZE_HEIGTH},
		&sdl.Rect{
			X: int32(p.Posx),
			Y:  int32(p.Posy),
			W: settings.PLAYER_SIZE_WIDTH,
			H: settings.PLAYER_SIZE_HEIGTH},
	)
} 