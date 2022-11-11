package configs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Coordinates struct{
	X int
	Y int
}

type MyImage struct {
	Img *ebiten.Image 
}

const WIDTH, HEIGHT = 8, 8 // in titles/squares
const PIXEL_LENGTH = 70
const START_LENGTH = 3
const NOT_MOVING = 0 
const UP = 2
const LEFT = 1
const RIGHT = 3 
const DOWN = 4
const QUIT_GAME = 5
const FPS = 10
