package main

import (
	"image/color"
	"log"
	"math/rand"
	"time"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

// IMPORT VARIABLES

const WIDTH, HEIGHT = 8, 8 // in titles/squares
const PIXEL_LENGTH = 70
const START_LENGTH = 3
const NOT_MOVING = 0 
const UP = 2
const LEFT = 1
const RIGHT = 3 
const DOWN = 4
const FPS = 8

var my_screenWidth = WIDTH*PIXEL_LENGTH
var my_screenHeight = HEIGHT*PIXEL_LENGTH
var board [WIDTH*HEIGHT]MyImage
var last_drawn time.Time 

var current_snake_length = START_LENGTH
var snake_pos []Coordinates
var apple_pos Coordinates
var r1 *rand.Rand
var direction = NOT_MOVING // 1:left 2:up 3:right 4:down

type Game struct{}

type Coordinates struct{
	X int
	Y int
}

type MyImage struct {
	img *ebiten.Image 
}


func get_opposite_direction(direction int) int {
	switch direction {
	case LEFT:
		return RIGHT	
	case RIGHT:
		return LEFT
	case UP:
		return DOWN
	case DOWN:
		return UP
	default:
		return direction
	}
}

func get_new_head_pos(d int) Coordinates {
	head_pos := snake_pos[len(snake_pos) - 1]
	switch d {
	case LEFT:
		return Coordinates{X: head_pos.X-1, Y: head_pos.Y}	
	case RIGHT:
		return Coordinates{X: head_pos.X+1, Y: head_pos.Y}	
	case UP:
		return Coordinates{X: head_pos.X, Y: head_pos.Y-1}	
	case DOWN:
		return Coordinates{X: head_pos.X, Y: head_pos.Y+1}	
	default:
		return Coordinates{X: head_pos.X, Y: head_pos.Y}	
	}
}

func adjust_out_of_board(pos Coordinates) Coordinates {
	if pos.X < 0 {
		pos.X = WIDTH - 1
	}
	if pos.X > WIDTH - 1 {
		pos.X = 0
	}
	if pos.Y < 0 {
		pos.Y = HEIGHT - 1
	}
	if pos.Y > HEIGHT - 1 {
		pos.Y = 0
	}
	return pos
}

func (g *Game) Update() error {
	new_direction := get_user_input()
	if new_direction == NOT_MOVING {
		return nil
	}
	if new_direction != get_opposite_direction(direction) {
		direction = new_direction
	}
	new_head_pos := get_new_head_pos(direction)

	// check collision:
	if has_coordinate(snake_pos, new_head_pos.X, new_head_pos.Y) {
		current_snake_length = START_LENGTH
	}
	new_head_pos = adjust_out_of_board(new_head_pos)

	snake_pos = append(snake_pos, new_head_pos)
	
	// apple check
	if new_head_pos.X == apple_pos.X && new_head_pos.Y == apple_pos.Y {
		current_snake_length += 1
		apple_pos = get_apple_pos()	
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
			tile.img.Clear()
			if has_coordinate(snake_pos, x, y) {
				tile.img.Fill(colornames.Green)
			} else if apple_pos.X == x && apple_pos.Y == y {
				tile.img.Fill(colornames.Red)
			} else {
				tile.img.Fill(color.Black)
			}
			op := &ebiten.DrawImageOptions{} 
			op.GeoM.Translate(float64(x*PIXEL_LENGTH), float64(y*PIXEL_LENGTH))
			// ebitenutil.DebugPrintAt(tile.img, strconv.Itoa(direction) , 0, 0)
			screen.DrawImage(tile.img, op)
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
	return direction
}

func has_coordinate(coors []Coordinates, x, y int) bool {
	for i := range coors {
		if coors[i].X == x && coors[i].Y == y {
			return true
		}
	}
	return false
}

func get_apple_pos() Coordinates {
	var x, y int
	x, y = rand.Intn(WIDTH), rand.Intn(HEIGHT)
	if has_coordinate(snake_pos, x, y) {
		return get_apple_pos()
	} else {
		return Coordinates{X: x, Y: y}
	}
}

func init_board() {
	for i := range board {
		board[i].img = ebiten.NewImage(PIXEL_LENGTH, PIXEL_LENGTH)
	}
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
	apple_pos = get_apple_pos()
	init_board()
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
