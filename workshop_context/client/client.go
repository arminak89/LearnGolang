package client

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Client struct {
	sleepTime time.Duration
	ctx       context.Context
}

type ClientOption func(*Client)

func WithContext(ctx context.Context) ClientOption {
	return func(c *Client) {
		c.ctx = context.WithValue(ctx, "ID", "123")
	}
}

func WithSleepTime(sleepTime time.Duration) ClientOption {
	return func(c *Client) {
		c.sleepTime = sleepTime
	}
}

// NewClient creates a default client
func NewClient(opts ...ClientOption) *Client {
	client := &Client{
		sleepTime: 1 * time.Second,
		ctx:       context.TODO(),
	}

	for _, opt := range opts {
		opt(client)
	}
	return client

}

func (c *Client) Sleep(format string, arg ...interface{}) error {
	select {
	case <-c.ctx.Done():
		return fmt.Errorf("mySleepContext[%s]: %v", c.ctx.Value("ID").(string), c.ctx.Err())
	case <-time.After(c.sleepTime):
		log.Printf(format, arg...)
		return nil
	}
}
