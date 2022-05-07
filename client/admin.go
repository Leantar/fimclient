package client

import (
	"context"
	"github.com/Leantar/fimproto/proto"
	"time"
)

func (c *Client) CreateAgentEndpoint(name string, watchedPaths []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	agent := proto.AgentEndpoint{
		Name:         name,
		WatchedPaths: watchedPaths,
	}

	_, err := c.client.CreateAgentEndpoint(ctx, &agent)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) CreateClientEndpoint(name string, roles []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := proto.ClientEndpoint{
		Name:  name,
		Roles: roles,
	}

	_, err := c.client.CreateClientEndpoint(ctx, &client)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteEndpoint(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	endpointName := proto.EndpointName{Name: name}

	_, err := c.client.DeleteEndpoint(ctx, &endpointName)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateEndpointWatchedPaths(name string, watchedPaths []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	agentUpdate := proto.AgentEndpoint{
		Name:         name,
		WatchedPaths: watchedPaths,
	}

	_, err := c.client.UpdateEndpointWatchedPaths(ctx, &agentUpdate)
	if err != nil {
		return err
	}

	return nil
}
