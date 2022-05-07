package client

import (
	"context"
	"github.com/Leantar/fimclient/models"
	"github.com/Leantar/fimproto/proto"
	"io"
	"time"
)

func (c *Client) GetAgents() ([]models.RecvAgent, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	empty := proto.Empty{}

	stream, err := c.client.GetAgents(ctx, &empty)
	if err != nil {
		return nil, err
	}

	agents := make([]models.RecvAgent, 0)

	for {
		agent, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		agents = append(agents, models.RecvAgent{
			Name:              agent.Name,
			HasBaseline:       agent.HasBaseline,
			BaselineIsCurrent: agent.BaselineIsCurrent,
			WatchedPaths:      agent.WatchedPaths,
		})

	}

	return agents, nil
}

func (c *Client) GetAlertsByAgent(name string) ([]models.RecvAlert, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	endpointName := proto.EndpointName{Name: name}

	stream, err := c.client.GetAlertsByAgent(ctx, &endpointName)
	if err != nil {
		return nil, err
	}

	alerts := make([]models.RecvAlert, 0)

	for {
		alert, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		alerts = append(alerts, models.RecvAlert{
			Kind:       alert.Kind,
			Difference: alert.Difference,
			Path:       alert.Path,
			IssuedAt:   alert.IssuedAt,
		})
	}

	return alerts, nil
}
