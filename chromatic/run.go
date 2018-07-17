// Copyright 2018 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package chromatic

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrorTimeoutExceeded = errors.New("timeout exceeded")
)

func RunWithConfig(config *Config) error {
	// Create a cancelable context. Cancelling this context will halt all
	// event, rpc, and exec machinery.
	cancelCtx, cancel := context.WithCancel(context.Background())

	// Create a timeout context. Waiting on this context will ensure we don't
	// hang forever.
	timedCtx, stop := context.WithTimeout(context.Background(), time.Duration(config.End.Timeout)*time.Second)

	// Create browser object with params
	browser := NewBrowser(cancelCtx, config.Start.URL, config.Browser.Flags...)

	// This will shutdown streams, rpc clients, and wait for the browser process to die.
	defer browser.Wait()
	defer stop()
	defer cancel()

	// Launch (exec) browser process
	if err := browser.Start(); err != nil {
		return err
	}

	// Greate remote debugger object
	client := NewClient(cancelCtx, browser.Address())

	// Browser may not have started listening yet
	time.Sleep(2 * time.Second)

	// Connect to remote debugger protocol RPC endpoint
	if err := client.Connect(); err != nil {
		//panic(err)
		return err
	}

	// Subscribe to load event stream
	eventChan, errChan, err := client.Events()
	if err != nil {
		return err
	}

	// Main "event loop"
	for {
		select {
		case <-timedCtx.Done():
			return ErrorTimeoutExceeded

		case err := <-errChan:
			return err

		case page := <-eventChan:
			if Match(page, config.End.Title, config.End.URL, config.End.Cookie) {
				output, _ := Report(page)
				fmt.Println(output)
				return nil
			}
		}
	}
}

func Run(filename string) error {
	// Parse config file
	config, err := Load(filename)
	if err != nil {
		return err
	}

	return RunWithConfig(config)
}
