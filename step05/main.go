package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
)

// Player is the player character \o/
type Player struct {
	row int
	col int
}

var player Player

// Ghost is the enemy that chases the player :O
type Ghost struct {
	row int
	col int
}

var ghosts []*Ghost

// START OMIT
func loadMaze() error {
	f, err := os.Open("maze01.txt") // OMIT
	if err != nil {                 // OMIT
		return err // OMIT
	} // OMIT
	defer f.Close() // OMIT
	// OMIT
	scanner := bufio.NewScanner(f) // OMIT
	for scanner.Scan() {           // OMIT
		line := scanner.Text()    // OMIT
		maze = append(maze, line) // OMIT
	} // OMIT
	// load maze omitted...

	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = Player{row, col}
			case 'G':
				ghosts = append(ghosts, &Ghost{row, col})
			case '.': // HL
				numDots++ // HL
			} // HL
		}
	}

	return nil
}

// END OMIT

var maze []string
var score int
var numDots int
var lives = 1

func clearScreen() {
	fmt.Printf("\x1b[2J")
	moveCursor(0, 0)
}

func moveCursor(row, col int) {
	fmt.Printf("\x1b[%d;%df", row+1, col+1)
}

// START2 OMIT
func printScreen() {
	clearScreen()
	for _, line := range maze {
		for _, chr := range line {
			switch chr {
			case '#': // HL
				fallthrough // HL
			case '.': // HL
				fmt.Printf("%c", chr) // HL
			default:
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	// OMIT
	moveCursor(player.row, player.col) // OMIT
	fmt.Printf("P")                    // OMIT
	// OMIT
	for _, g := range ghosts { // OMIT
		moveCursor(g.row, g.col) // OMIT
		fmt.Printf("G")          // OMIT
	} // OMIT
	// OMIT
	moveCursor(len(maze)+1, 0)                         // OMIT
	fmt.Printf("Score: %v\tLives: %v\n", score, lives) // OMIT
	// lines omitted...
}

// END2 OMIT

func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 {
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}

	return "", nil
}

func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol

	switch dir {
	case "UP":
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(maze)-1 {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(maze[0]) - 1
		}
	}

	if maze[newRow][newCol] == '#' {
		newRow = oldRow
		newCol = oldCol
	}

	return
}

func movePlayer(dir string) {
	player.row, player.col = makeMove(player.row, player.col, dir)
	switch maze[player.row][player.col] { // HL
	case '.': // HL
		numDots-- // HL
		score++   // HL
		// Remove dot from the maze
		maze[player.row] = maze[player.row][0:player.col] + " " + maze[player.row][player.col+1:] // HL
	} // HL
}

func drawDirection() string {
	dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "RIGHT",
		3: "LEFT",
	}
	return move[dir]
}

func moveGhosts() {
	for _, g := range ghosts {
		dir := drawDirection()
		g.row, g.col = makeMove(g.row, g.col, dir)
	}
}

func init() {
	cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode terminal: %v\n", err)
	}
}

func cleanup() {
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cooked mode terminal: %v\n", err)
	}
}

func main() {
	// initialize game
	defer cleanup()

	// load resources
	err := loadMaze()
	if err != nil {
		log.Printf("Error loading maze: %v\n", err)
		return
	}

	// game loop
	for {
		// update screen
		printScreen()

		// process input
		input, err := readInput()
		if err != nil {
			log.Printf("Error reading input: %v", err)
			break
		}

		// process movement
		movePlayer(input)
		moveGhosts()

		// process collisions
		for _, g := range ghosts { // HL
			if player.row == g.row && player.col == g.col { // HL
				lives = 0 // HL
			} // HL
		} // HL

		// check game over
		if input == "ESC" || numDots == 0 || lives == 0 { // HL
			break
		}

		// repeat
	}
}
