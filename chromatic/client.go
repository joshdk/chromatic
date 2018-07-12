package chromatic

import (
	"context"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/rpcc"
)

// NewClient creates Client object that can be used to connect to and receive
// events from the Chrome Debugging Protocol.
func NewClient(ctx context.Context, url string) *Client {
	return &Client{
		ctx: ctx,
		url: url,
	}
}

// Client represents a single connection to the Chrome Debugging Protocol.
type Client struct {
	ctx    context.Context
	url    string
	client *cdp.Client
	id     string
}

// Connect to the Chrome Debugging Protocol endpoint.
func (c *Client) Connect() error {
	// Contact the Dev Tools JSON API and retrieve the current page target.
	pageTarget, err := devtool.New(c.url).Get(c.ctx, devtool.Page)
	if err != nil {
		return err
	}

	// Create an RPC connection to the Chrome Debugging Protocol websocket
	// endpoint.
	conn, err := rpcc.DialContext(c.ctx, pageTarget.WebSocketDebuggerURL)
	if err != nil {
		return err
	}

	c.id = pageTarget.ID
	c.client = cdp.NewClient(conn)
	return nil
}

// Events listens for page load events and returns any errors on a channel.
func (c *Client) Events() (<-chan error, error) {
	var errors = make(chan error)

	// Create a client to buffer LoadEventFired events.
	events, err := c.client.Page.LoadEventFired(c.ctx)
	if err != nil {
		return nil, err
	}

	// Enable events on the Page domain.
	if err = c.client.Page.Enable(c.ctx); err != nil {
		return nil, err
	}

	go func() {
		defer events.Close()

		for {
			// Wait until we have a LoadEventFired event.
			if _, err := events.Recv(); err != nil {
				errors <- err
				return
			}
		}
	}()

	return errors, nil
}
