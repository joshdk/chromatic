// Copyright 2018 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

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
		switch err {
		case chromatic.ErrorTimeoutExceeded:
			os.Exit(2)
		default:
			os.Exit(1)
		}
	}
}
