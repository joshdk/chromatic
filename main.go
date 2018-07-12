package main

import (
	"fmt"
	"os"

	"github.com/joshdk/chromatic/chromatic"
)

func mainFunc(args []string) error {
	return chromatic.Run("chromatic.yml")
}

func main() {
	if err := mainFunc(nil); err != nil {
		fmt.Fprintf(os.Stderr, "chromatic: %s\n", err.Error())
		os.Exit(1)
	}
}
