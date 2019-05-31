package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/danicat/pacgo/stepxx/config"
	"github.com/danicat/pacgo/stepxx/maze"
)

var (
	configFile = flag.String("config-file", "config.json", "path to custom configuration file")
	mazeFile   = flag.String("maze-file", "maze01.txt", "path to a custom maze file")
)

func main() {
	flag.Parse()

	// initialize game
	initialize()
	defer cleanup()

	// load resources
	err := config.Load(*configFile)
	if err != nil {
		log.Printf("Error loading configuration: %v\n", err)
		return
	}

	err = maze.Load(*mazeFile)
	if err != nil {
		log.Printf("Error loading maze: %v\n", err)
		return
	}

	// game loop
	for {
		// process movement
		for _, s := range sprites {
			go s.Move()
		}

		// update screen
		printScreen()

		// check game over
		if numDots == 0 || player.lives == 0 {
			if player.lives == 0 {
				moveCursor(player.Pos())
				fmt.Printf(config.Death())
				moveCursor(len(maze)+2, 0)
			}
			break
		}

		// wait before rendering next frame
		time.Sleep(1000 / config.FrameRate() * time.Millisecond)
	}
}
