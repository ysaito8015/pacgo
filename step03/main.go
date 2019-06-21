package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Player is the player character \o/
type Player struct {
	row int
	col int
}

var player Player

// START1 OMIT
func loadMaze() error {
	// file open omitted...
	f, err := os.Open("maze01.txt") // OMIT
	if err != nil {                 // OMIT
		return err // OMIT
	} // OMIT
	defer f.Close() // OMIT

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	for row, line := range maze { // HL
		for col, char := range line { // HL
			switch char { // HL
			case 'P': // HL
				player = Player{row, col} // HL
			} // HL
		} // HL
	} // HL

	return nil
}

// END1 OMIT

var maze []string

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
	for _, line := range maze { // HL
		for _, chr := range line { // HL
			switch chr { // HL
			case '#': // HL
				fmt.Printf("%c", chr) // HL
			default: // HL
				fmt.Printf(" ") // HL
			} // HL
		} // HL
		fmt.Printf("\n") // HL
	} // HL

	moveCursor(player.row, player.col) // HL
	fmt.Printf("P")                    // HL
	// OMIT
	moveCursor(len(maze)+1, 0)                          // OMIT
	fmt.Printf("Row %v Col %v", player.row, player.col) // OMIT
}

// END2 OMIT

// START3 OMIT
func readInput() (string, error) {
	buffer := make([]byte, 100) // OMIT
	// OMIT
	cnt, err := os.Stdin.Read(buffer) // OMIT
	if err != nil {                   // OMIT
		return "", err // OMIT
	} // OMIT
	// lines omitted...

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 { // HL
		if buffer[0] == 0x1b && buffer[1] == '[' { // HL
			switch buffer[2] { // HL
			case 'A': // HL
				return "UP", nil // HL
			case 'B': // HL
				return "DOWN", nil // HL
			case 'C': // HL
				return "RIGHT", nil // HL
			case 'D': // HL
				return "LEFT", nil // HL
			} // HL
		} // HL
	} // HL

	return "", nil
}

// END3 OMIT

func makeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol // HL

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
		// omitted similar cases for RIGHT and LEFT...
	case "RIGHT": // OMIT
		newCol = newCol + 1         // OMIT
		if newCol == len(maze[0]) { // OMIT
			newCol = 0 // OMIT
		} // OMIT
	case "LEFT": // OMIT
		newCol = newCol - 1 // OMIT
		if newCol < 0 {     // OMIT
			newCol = len(maze[0]) - 1 // OMIT
		} // OMIT
	}

	if maze[newRow][newCol] == '#' {
		newRow = oldRow
		newCol = oldCol
	}

	return // HL
}

func movePlayer(dir string) {
	player.row, player.col = makeMove(player.row, player.col, dir)
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

	// START4 OMIT
	// game loop
	for {
		// update screen
		printScreen()

		// process input (omitted)
		input, err := readInput() // OMIT
		if err != nil {           // OMIT
			log.Printf("Error reading input: %v", err) // OMIT
			break                                      // OMIT
		} // OMIT

		// process movement
		movePlayer(input) // HL

		// process collisions

		// check game over
		if input == "ESC" {
			break
		}

		// repeat
	}
	// END4 OMIT
}
