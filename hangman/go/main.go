package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Coordinates struct {
	X int
	Y int
}

var hangman_img = [8][7]rune{
	{0 ,'\u213c','-','-','-','-',0},
	{0 , '|', 0, 0,'|', 0, 0,},
	{0 , '|', 0, 0,'O', 0, 0,},
	{0 , '|',0,'/','|','\\',0},
	{0 , '|', 0,0, '|',  0, 0},
	{0 , '|', 0,'/',0,'\\', 0},
	{0 , '|', 0,0,0,0, 0},
	{0,'\u22a5', 0,0,0,0, 0},
}

var words_slice = []string{"god of war", "breaking bad", "silent hill"}
var word string
var guessed_letters map[string]bool
const max_guesses int = 8
var man_coordinates []Coordinates

func init_state() {
	guessed_letters = make(map[string]bool) // reinit
	man_coordinates = []Coordinates{
		{5,5},{3,5},{4,4},{5,3},{3,3},{4,3},{4,2},
	}
	rand.Seed(time.Now().Unix())
	word = words_slice[rand.Intn(len(words_slice))]
}

func draw_hangman() {
	w,h := 7,8
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			val := hangman_img[y][x]
			if val == 0 {
				fmt.Print(" ")
			} else {
				cor := Coordinates{X: x, Y: y}
				print_val := string(val)
				for i := range man_coordinates {
					if man_coordinates[i] == cor {
						print_val = " "
					}
				}
				fmt.Print(print_val)
			}
		}
		fmt.Println("") // new line
	}
}

func draw_words() {
	for i := range word {
		fmt.Print(" ")
		if guessed_letters[string(word[i])] {
			fmt.Print("\u0332" + string(word[i]))
		}else {
			if word[i] == ' ' {
				fmt.Print(string(word[i]))
			} else {
				fmt.Print("_")
			}
		}
	}
		fmt.Print(" ")
}

func letters_set(word string) map[rune]bool {
	res := make(map[rune]bool) 

	for i := range word {
		if word[i] != ' '  {
			res[rune(word[i])] = true
		}
	}
	return res
}

func get_win_or_lose() string {
	if len(man_coordinates) == 0 {
		return "lose"
	}
	if len(guessed_letters) == len(letters_set(word)) {
		return "win"
	}
	return ""
}

func ask(x string) bool {
	// answer y or n only
	fmt.Print(x +": ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	for !(input == "y" || input == "n") {
		fmt.Println("Must type y or n!. Please type again:")
		scanner.Scan()
		input = strings.TrimSpace(scanner.Text())
	} 
		
	if input == "y" {
		return true
	} else {
		return false
	}
}

func clear_terminal() {
	clear_cmd := exec.Command("clear")
	clear_cmd.Stdout = os.Stdout
	clear_cmd.Run()
}

func main() {
	init_state()

	for {
		clear_terminal()
		draw_hangman()
		draw_words()

		win_or_lose := get_win_or_lose()

		if win_or_lose != "" {
			if win_or_lose == "win" {
				fmt.Println("\nyou won!")
			} else {
				fmt.Println("\nyou lost!")
			}

			again := ask("play again?")
			if again {
				init_state()
				continue
			} else {
				os.Exit(0)
			}
		} 
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter valid English char: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		for guessed_letters[input] == true || len(input) == 0{
			fmt.Println("Existing letter or empty input!")
			fmt.Print("Enter another letter: ")
			scanner.Scan()
			input = strings.TrimSpace(scanner.Text())
		} 
		input = input[:1]
		if strings.Contains(word, input) {
			guessed_letters[input] = true
		} else {
			man_coordinates = man_coordinates[:len(man_coordinates)-1]
		}
	}
}