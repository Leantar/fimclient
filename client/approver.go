package client

import (
	"context"
	"github.com/Leantar/fimproto/proto"
	"time"
)

func (c *Client) CreateBaselineUpdateApproval(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	endpointName := proto.EndpointName{Name: name}

	_, err := c.client.CreateBaselineUpdateApproval(ctx, &endpointName)
	if err != nil {
		return err
	}
	return nil
}
