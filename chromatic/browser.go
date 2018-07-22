// Copyright 2018 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package chromatic

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	defaultChromeName = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	defaultChromeArgs = []string{
		"--no-experiments",
		"--no-first-run",
		"--remote-debugging-port=9222",
	}
	defaultDebuggerAddress = "http://127.0.0.1:9222"
)

func tempDir() (string, func(), error) {
	name, err := ioutil.TempDir("", "chromatic-")
	if err != nil {
		return "", nil, err
	}

	return name, func() {
		if err := os.RemoveAll(name); err != nil {
			panic(err)
		}
	}, nil
}

// NewBrowser creates a Browser object that can start Chrome with optional
// args, and will initially load the given url.
func NewBrowser(ctx context.Context, url string, args ...string) (*Browser, error) {
	var (
		dataDir, cleanup, err = tempDir()
		b                     = Browser{
			name:    defaultChromeName,
			args:    defaultChromeArgs,
			cleanup: cleanup,
		}
	)

	if err != nil {
		return nil, err
	}

	b.args = append(b.args, "--user-data-dir="+dataDir)
	b.args = append(b.args, args...)
	b.args = append(b.args, url)

	fmt.Printf("args are: %+v\n", b.args)

	b.ctx, b.cancel = context.WithCancel(ctx)
	b.cmd = exec.CommandContext(b.ctx, b.name, b.args...)
	return &b, nil
}

// Browser represents a single Chrome instance that can be started, stopped,
// and waited on.
type Browser struct {
	name    string
	args    []string
	ctx     context.Context
	cmd     *exec.Cmd
	cancel  context.CancelFunc
	cleanup func()
}

// Address returns the Chrome Debugging Protocol address.
func (b *Browser) Address() string {
	return defaultDebuggerAddress
}

// Start creates a Chrome process, and will make a browser window appear.
func (b *Browser) Start() error {
	return b.cmd.Start()
}

// Stop terminates the Chrome process.
func (b *Browser) Stop() error {
	b.cancel()
	return b.Wait()
}

// Wait returns when the Chrome process eventually terminates.
func (b *Browser) Wait() error {
	defer b.cleanup()
	return b.cmd.Wait()
}
