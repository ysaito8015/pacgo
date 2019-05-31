package screen

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/danicat/pacgo/stepxx/config"
)

func Init() {
	cbTerm := exec.Command("/bin/stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cbreak mode terminal: %v\n", err)
	}
}

// /bin/stty -cbreak echo
func Cleanup() {
	cookedTerm := exec.Command("/bin/stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("Unable to activate cooked mode terminal: %v\n", err)
	}
}

func Clear() {
	fmt.Printf("\x1b[2J")
	SetCursor(0, 0)
}

func SetCursor(row, col int) {
	if config.UseEmoji() {
		fmt.Printf("\x1b[%d;%df", row+1, col*2+1)
	} else {
		fmt.Printf("\x1b[%d;%df", row+1, col+1)
	}
}
