package cmd

import (
	"fmt"
	"github.com/Leantar/fimclient/client"
	"github.com/Leantar/fimclient/models"
	"github.com/Leantar/fimclient/modules/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

func printAgents(_ *cobra.Command, _ []string) error {
	var conf client.Config
	err := config.FromYamlFile(confPath, &conf)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	c, err := client.NewConnectedClient(conf)
	if err != nil {
		return err
	}

	agents, err := c.GetAgents()
	if err != nil {
		return err
	}

	for _, agent := range agents {
		err := yaml.NewEncoder(os.Stdout).Encode(&agent)
		if err != nil {
			return err
		}
		fmt.Printf("---------------------------\n")
	}

	return nil
}

func printAlerts(_ *cobra.Command, _ []string) error {
	var conf client.Config
	err := config.FromYamlFile(confPath, &conf)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	c, err := client.NewConnectedClient(conf)
	if err != nil {
		return err
	}

	alerts, err := c.GetAlertsByAgent(agentName)
	if err != nil {
		return err
	}

	for _, alert := range alerts {
		al := models.PrintableAlert{
			Kind:       alert.Kind,
			Difference: alert.Difference,
			Path:       alert.Path,
			IssuedAt:   time.Unix(alert.IssuedAt, 0).Format(time.UnixDate),
		}

		err := yaml.NewEncoder(os.Stdout).Encode(&al)
		if err != nil {
			return err
		}
		fmt.Printf("---------------------------\n")
	}

	return nil
}
