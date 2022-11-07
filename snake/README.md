## rules
- snake dies when it hits itself
- gets apple, length += 1
- if a part of its body moves out of board, it appears in the opposite side of the board
- size of board is 10x10

## pseudo code

```
Coordinates struct {
  x int
  y int
}


// important global variables
width, height = 10, 10
pixel_length = 15
board = [width*height]Coordinates

current_snake_length = 3
snake_pos = a queue of coordinates []Coordinates
apple_pos = get_apple_pos() // get a random pos on board (not on snake's body)
direction = 0 (1: left; 2: up; 3; right; 4; down)


// game loop
while True:
  1. get user input --
  // arrow keys
  new_direction = get_user_input()
  
  // hanle case opposite of current direction => same direction
  if new_direction != opposite of direction:
    direction = new_direction
 
  2. process user input + compute new state --
  // based on arrow key => new snake head position
  new_head_pos = get_new_head_pos(new_direction)
  
  // check losing condition:
  if new_head_pos in snake_pos:
    lose()
  else:
    snake_pos.push(new_head_pos)
    
  // check if on an apple:
  if new_head_pos == apple_pos:
    current_snake_length += 1
  
  // make sure snake can not exceed current length
  while len(snake_pos) > current_snake_length:
    // pop the queue
    snake_pos = snake_pos[1:]
  
  3. draw black board + snake + apple
  draw()
```
