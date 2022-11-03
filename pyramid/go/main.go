package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func pyramid(h, w int) [][]int {
	arr := make([][]int, h)
	for y := 0; y < h; y++ {
		arr[y] = make([]int, w) 
		for x := 0; x < w; x++ {
			if y <= x && x <= (w-1-y) {
				arr[y][x] = 1
			}
		}
	}
	return arr
}

func draw(arr [][]int, reverse bool) {
	if reverse {
		for _, v_y := range arr {
			for _, v_x := range v_y {
				v_to_print := " "
				if v_x == 1 {
					v_to_print = "*"	
				}
				fmt.Printf(v_to_print)
			}
			fmt.Printf("\n") // new line
		}
	} else {
		for y := len(arr) - 1; y >= 0; y-- {
			for x := 0; x < len(arr[y]); x++ {
				v_to_print := " "
				if arr[y][x] == 1 {
					v_to_print = "*"	
				}
				fmt.Printf(v_to_print)
			}
			fmt.Printf("\n") // new line
		}

	}

}

func main() {
	// get user input: height && reverse
	var height int
	var reverse bool = false
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type your height(int): ")
	scanner.Scan()
	input, err := strconv.ParseInt(scanner.Text(), 10, 64)
	for err != nil {
		fmt.Println("Please enter an int!")
		fmt.Println("Type your height(int): ")
		scanner.Scan()
		input, err = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	height = int(input)
	width := height * 2 - 1

	fmt.Println("Do you want it reverse? (Type y or n)")
	scanner.Scan()
	reverse_input := scanner.Text()	
	for ! (reverse_input == "y" || reverse_input == "n") {
		fmt.Println("Do you want it reverse? (Type y or n)")
		scanner.Scan()
		reverse_input = scanner.Text()	
	}
	if reverse_input == "y" {
		reverse = true
	}
	p := pyramid(height, width)
	draw(p, reverse)
}
