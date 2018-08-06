// Copyright 2018 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package main

import (
	"fmt"
	"os"

	"flag"

	"github.com/joshdk/chromatic/chromatic"
)

var (
	version = "development"
)

func main() {
	versionFlag := flag.Bool("version", false, "Displays the version and exits.")
	flag.Parse()

	if *versionFlag == true {
		fmt.Println(version)
		os.Exit(2)

	} else if err := mainFunc(flag.Arg(0)); err != nil {
		fmt.Fprintf(os.Stderr, "chromatic: %s\n", err.Error())
		switch err {
		case chromatic.ErrorTimeoutExceeded:
			os.Exit(2)
		default:
			os.Exit(1)
		}
	}
}

func mainFunc(configFile string) error {
	if configFile == "" {
		configFile = "chromatic.yml"
	}
	return chromatic.Run(configFile)
}
