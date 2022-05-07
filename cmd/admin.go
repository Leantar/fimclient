package cmd

import (
	"fmt"
	"github.com/Leantar/fimclient/client"
	"github.com/Leantar/fimclient/models"
	"github.com/Leantar/fimclient/modules/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

func deleteEndpoint(_ *cobra.Command, _ []string) error {
	var conf client.Config
	err := config.FromYamlFile(confPath, &conf)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	c, err := client.NewConnectedClient(conf)
	if err != nil {
		return err
	}

	return c.DeleteEndpoint(agentName)
}

func createAgent(_ *cobra.Command, _ []string) error {
	var conf client.Config
	err := config.FromYamlFile(confPath, &conf)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	c, err := client.NewConnectedClient(conf)
	if err != nil {
		return err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	var agents []models.Agent
	err = yaml.NewDecoder(f).Decode(&agents)
	if err != nil {
		return err
	}

	for _, ag := range agents {
		err = c.CreateAgentEndpoint(ag.Name, ag.WatchedPaths)
		if err != nil {
			return err
		}
	}

	return nil
}

func createClient(_ *cobra.Command, _ []string) error {
	var conf client.Config
	err := config.FromYamlFile(confPath, &conf)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	c, err := client.NewConnectedClient(conf)
	if err != nil {
		return err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	var clients []models.Client
	err = yaml.NewDecoder(f).Decode(&clients)
	if err != nil {
		return err
	}

	for _, cl := range clients {
		err = c.CreateClientEndpoint(cl.Name, cl.Roles)
		if err != nil {
			return err
		}
	}

	return nil
}

func updateWatchedPaths(_ *cobra.Command, _ []string) error {
	var conf client.Config
	err := config.FromYamlFile(confPath, &conf)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	c, err := client.NewConnectedClient(conf)
	if err != nil {
		return err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	var update models.WatchedPathsUpdate
	err = yaml.NewDecoder(f).Decode(&update)
	if err != nil {
		return err
	}

	return c.UpdateEndpointWatchedPaths(update.Name, update.WatchedPaths)
}
