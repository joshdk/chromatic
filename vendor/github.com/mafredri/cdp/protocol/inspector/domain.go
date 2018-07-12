// Code generated by cdpgen. DO NOT EDIT.

// Package inspector implements the Inspector domain.
package inspector

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Inspector domain.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Inspector domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Disable invokes the Inspector method. Disables inspector domain
// notifications.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Inspector.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Inspector", Op: "Disable", Err: err}
	}
	return
}

// Enable invokes the Inspector method. Enables inspector domain
// notifications.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Inspector.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Inspector", Op: "Enable", Err: err}
	}
	return
}

func (d *domainClient) Detached(ctx context.Context) (DetachedClient, error) {
	s, err := rpcc.NewStream(ctx, "Inspector.detached", d.conn)
	if err != nil {
		return nil, err
	}
	return &detachedClient{Stream: s}, nil
}

type detachedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *detachedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *detachedClient) Recv() (*DetachedReply, error) {
	event := new(DetachedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Inspector", Op: "Detached Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) TargetCrashed(ctx context.Context) (TargetCrashedClient, error) {
	s, err := rpcc.NewStream(ctx, "Inspector.targetCrashed", d.conn)
	if err != nil {
		return nil, err
	}
	return &targetCrashedClient{Stream: s}, nil
}

type targetCrashedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *targetCrashedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *targetCrashedClient) Recv() (*TargetCrashedReply, error) {
	event := new(TargetCrashedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Inspector", Op: "TargetCrashed Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) TargetReloadedAfterCrash(ctx context.Context) (TargetReloadedAfterCrashClient, error) {
	s, err := rpcc.NewStream(ctx, "Inspector.targetReloadedAfterCrash", d.conn)
	if err != nil {
		return nil, err
	}
	return &targetReloadedAfterCrashClient{Stream: s}, nil
}

type targetReloadedAfterCrashClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *targetReloadedAfterCrashClient) GetStream() rpcc.Stream { return c.Stream }

func (c *targetReloadedAfterCrashClient) Recv() (*TargetReloadedAfterCrashReply, error) {
	event := new(TargetReloadedAfterCrashReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Inspector", Op: "TargetReloadedAfterCrash Recv", Err: err}
	}
	return event, nil
}
