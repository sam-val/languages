package main

import (
	"fmt"
	"time"
)

const FPS = 2
const WIDTH = 10
const HEIGHT = 10


func count_neighbours(state [][]int, x, y int) int {
	exists := func(state [][]int, x, y int) bool {
		if x < 0 || x >= len(state[0]) {
			return false
		}
		if y < 0 || y >= len(state) {
			return false
		}
		if state[y][x] == 1 {
			return true
		}
		
		return false
	}
	neighs := 0
	neighbour_cors := []struct{X,Y int}{
		{x-1, y-1},
		{x-1, y+1},
		{x-1, y},
		{x+1, y-1},
		{x+1, y+1},
		{x+1, y},
		{x, y-1},
		{x, y+1},
	}

	for _, cor := range neighbour_cors {
		if exists(state, cor.X, cor.Y) {
			neighs++
		}
	}
	return neighs
}

func compute(state [][]int, w, h int) [][]int {
	// make empty 2d array
	new_state := make([][]int, h)
	for i := 0; i<h;i++ {
		new_state[i] = make([]int, w)
	}

	var neighbours int
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			neighbours = count_neighbours(state, j, i)
			if state[i][j] == 1 {
				if neighbours == 2 || neighbours == 3 {
					new_state[i][j] = 1
				}
			} else {
				if neighbours == 3 {
					new_state[i][j] = 1
				}
			}
		}
	}
	return new_state
}

func draw(state [][]int, w, h int) {
	fmt.Println("")
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if state[i][j] == 1 {
				fmt.Print("#|")
			} else {
				fmt.Print(" |")
			}
		}
		fmt.Println()
	} 
}

var current_state = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 0, 0, 0, 0, 0, 0, 0},
	{0, 1, 1, 0, 0, 0, 0, 1, 1, 1},
	{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 1, 0, 0, 0, 0, 0},

}

func main() {
	for {
		new_state := compute(current_state, WIDTH, HEIGHT)

		// draw
		draw(new_state, WIDTH, HEIGHT)

		current_state = new_state

		time.Sleep((1000/FPS)*time.Millisecond)
	}
}