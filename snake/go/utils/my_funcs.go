package utils

import (
	"math/rand"

	. "github.com/sam-val/languages/tree/main/snake/go/configs"
)

func GetOppositeDirection(direction int) int {
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

func GetNewHeadPos(snake_pos []Coordinates, d int) Coordinates {
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

func AdjustOutOfBoard(pos Coordinates) Coordinates {
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

func HasCoordinate(coors []Coordinates, x, y int) bool {
	for i := range coors {
		if coors[i].X == x && coors[i].Y == y {
			return true
		}
	}
	return false
}

func GetApplePos(snake_pos []Coordinates) Coordinates {
	var x, y int
	x, y = rand.Intn(WIDTH), rand.Intn(HEIGHT)
	if HasCoordinate(snake_pos, x, y) {
		return GetApplePos(snake_pos)
	} else {
		return Coordinates{X: x, Y: y}
	}
}
