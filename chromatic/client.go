package chromatic

import (
	"context"

	"github.com/mafredri/cdp"
	"github.com/mafredri/cdp/devtool"
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/target"
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
func (c *Client) Events() (<-chan Page, <-chan error, error) {
	var (
		errors = make(chan error)
		pages  = make(chan Page)
	)

	// Create a client to buffer LoadEventFired events.
	events, err := c.client.Page.LoadEventFired(c.ctx)
	if err != nil {
		return nil, nil, err
	}

	// Enable events on the Page domain.
	if err = c.client.Page.Enable(c.ctx); err != nil {
		return nil, nil, err
	}

	go func() {
		defer events.Close()

		for {
			// Wait until we have a LoadEventFired event.
			if _, err := events.Recv(); err != nil {
				errors <- err
				return
			}

			// Extract some useful information from the page
			page, err := pageInfo(c.ctx, c.client, c.id)
			if err != nil {
				errors <- err
				return
			}

			pages <- *page
		}
	}()

	return pages, errors, nil
}

func pageInfo(ctx context.Context, client *cdp.Client, tid string) (*Page, error) {
	var id = target.ID(tid)

	targetinfo, err := client.Target.GetTargetInfo(ctx, &target.GetTargetInfoArgs{TargetID: &id})
	if err != nil {
		return nil, err
	}

	cookies, err := client.Network.GetAllCookies(ctx)
	if err != nil {
		return nil, err
	}

	doc, err := client.DOM.GetDocument(ctx, nil)
	if err != nil {
		return nil, err
	}

	result, err := client.DOM.GetOuterHTML(ctx, &dom.GetOuterHTMLArgs{
		NodeID: &doc.Root.NodeID,
	})
	if err != nil {
		return nil, err
	}

	return &Page{
		URL:     targetinfo.TargetInfo.URL,
		Title:   targetinfo.TargetInfo.Title,
		Cookies: cookies.Cookies,
		Body:    result.OuterHTML,
	}, nil
}
