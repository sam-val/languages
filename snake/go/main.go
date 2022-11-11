package main

import (
	"errors"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"

	. "github.com/sam-val/languages/tree/main/snake/go/configs"
	. "github.com/sam-val/languages/tree/main/snake/go/utils"
)

type Game struct{}

var my_screenWidth = WIDTH*PIXEL_LENGTH
var my_screenHeight = HEIGHT*PIXEL_LENGTH
var board [WIDTH*HEIGHT]MyImage
var last_drawn time.Time 

var current_snake_length = START_LENGTH
var snake_pos []Coordinates
var apple_pos Coordinates
var r1 *rand.Rand
var direction = NOT_MOVING // 1:left 2:up 3:right 4:down
var Quit = errors.New("quit")

func (g *Game) Update() error {
	input := get_user_input()
	if input == QUIT_GAME {
		return Quit
	}
	if input == NOT_MOVING {
		return nil
	}
	if input != GetOppositeDirection(direction) {
		direction = input
	}
	new_head_pos := GetNewHeadPos(snake_pos, direction)

	// check collision:
	if HasCoordinate(snake_pos, new_head_pos.X, new_head_pos.Y) {
		current_snake_length = START_LENGTH
	}
	new_head_pos = AdjustOutOfBoard(new_head_pos)

	snake_pos = append(snake_pos, new_head_pos)
	
	// apple check
	if new_head_pos.X == apple_pos.X && new_head_pos.Y == apple_pos.Y {
		current_snake_length += 1
		apple_pos = GetApplePos(snake_pos)	
	}

	for len(snake_pos) > current_snake_length {
		snake_pos = snake_pos[1:]
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			pos := WIDTH*y+x
			tile := board[pos]
			tile.Img.Clear()
			if HasCoordinate(snake_pos, x, y) {
				tile.Img.Fill(colornames.Green)
			} else if apple_pos.X == x && apple_pos.Y == y {
				tile.Img.Fill(colornames.Red)
			} else {
				tile.Img.Fill(color.Black)
			}
			op := &ebiten.DrawImageOptions{} 
			op.GeoM.Translate(float64(x*PIXEL_LENGTH), float64(y*PIXEL_LENGTH))
			// ebitenutil.DebugPrintAt(tile.Img, strconv.Itoa(direction) , 0, 0)
			screen.DrawImage(tile.Img, op)
		}	
	}
	last_drawn = time.Now()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return my_screenWidth, my_screenHeight
}

func get_user_input() int {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		return UP
	} 
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		return DOWN
	} 
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		return LEFT
	} 
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		return RIGHT
	} 
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return QUIT_GAME
	} 
	return direction
}


func init() {
	// prepare rand seed
	rand.Seed(time.Now().UnixNano())
}

func main() {
	ebiten.SetWindowSize(my_screenWidth, my_screenHeight)
	ebiten.SetWindowTitle("Snake")
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetTPS(FPS)
	snake_pos = append(snake_pos, Coordinates{X: WIDTH/2, Y: HEIGHT/2})
	apple_pos = GetApplePos(snake_pos)
	init_board := func() {
		for i := range board {
			board[i].Img = ebiten.NewImage(PIXEL_LENGTH, PIXEL_LENGTH)
		}
	}
	init_board()
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		if err == Quit {
			return
		}
		log.Fatal(err)
	}
}
