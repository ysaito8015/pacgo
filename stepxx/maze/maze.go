package maze

import (
	"bufio"
	"fmt"
	"os"

	"github.com/danicat/pacgo/stepxx/config"
	"github.com/danicat/pacgo/stepxx/screen"
	"github.com/danicat/pacgo/stepxx/sprite"
)

// Maze holds the game information
type Maze struct {
	layout  []string
	dots    int
	Player  *sprite.Player
	Sprites []sprite.Sprite
}

// Load the game maze
func Load(file string) (*Maze, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var maze Maze
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze.layout = append(maze.layout, line)
	}

	if !maze.Validate() {
		return nil, fmt.Errorf("invalid maze format")
	}

	for row, line := range maze.layout {
		for col, char := range line {
			switch char {
			case 'P':
				player = sprite.NewPlayer(row, col, config.Lives(), config.Player())
				sprites = append(sprites, player)
			case 'G':
				sprites = append(sprites, sprite.NewGhost(row, col, config.Ghost()))
			case 'C':
				sprites = append(sprites, sprite.NewChaser(row, col, config.Chaser()))
			case '.':
				maze.dots++
			}
		}
	}

	return &maze, nil
}

// PrintFrame clear the screen and render a new frame
func (m *Maze) PrintFrame() {
	screen.Clear()
	for _, line := range m.layout {
		for _, chr := range line {
			switch chr {
			case '#':
				fmt.Printf(config.Wall())
			case '.':
				fmt.Printf(config.Dot())
			default:
				fmt.Printf(config.Space())
			}
		}
		fmt.Printf("\n")
	}

	for _, s := range sprites {
		screen.SetCursor(s.Pos())
		fmt.Printf(s.Img())
	}

	screen.SetCursor(len(m.layout)+1, 0)
	fmt.Printf("Score: %v\tLives: %v\n", m.Player.score, m.Player.lives)
}

func (m Maze) Rows() int {
	return len(m.layout)
}

func (m Maze) Cols() int {
	return len(m.layout[0])
}

// Validate returns true if the maze is valid
func (m Maze) Validate() bool {
	cols := len(m.layout[0])
	for _, line := range m.layout[1:] {
		if len(line) != cols {
			return false
		}
	}
	return true
}

// IsValid return true if row, col is a legal move
func (m Maze) IsValid(row, col int) bool {
	return row >= 0 && row < m.Rows() &&
		col >= 0 && col < m.Cols() &&
		m.layout[row][col] != "#"
}
