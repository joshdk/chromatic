// Copyright 2018 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package chromatic

import (
	"context"
	"os/exec"
)

var (
	defaultChromeName = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
	defaultChromeArgs = []string{
		"--no-experiments",
		"--no-first-run",
		"--remote-debugging-port=9222",
		"--user-data-dir=./tmp",
	}
	defaultDebuggerAddress = "http://127.0.0.1:9222"
)

// NewBrowser creates a Browser object that can start Chrome with optional
// args, and will initially load the given url.
func NewBrowser(ctx context.Context, url string, args ...string) *Browser {
	var b = Browser{
		name: defaultChromeName,
		args: defaultChromeArgs,
	}

	for _, arg := range args {
		b.args = append(b.args, arg)
	}
	b.args = append(b.args, url)

	b.ctx, b.cancel = context.WithCancel(ctx)
	b.cmd = exec.CommandContext(b.ctx, b.name, b.args...)
	return &b
}

// Browser represents a single Chrome instance that can be started, stopped,
// and waited on.
type Browser struct {
	name   string
	args   []string
	ctx    context.Context
	cmd    *exec.Cmd
	cancel context.CancelFunc
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
	return b.cmd.Wait()
}
