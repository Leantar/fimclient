package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var (
	confPath  string
	agentName string
	filePath  string
	rootCmd   = &cobra.Command{
		Use:   "fimclient",
		Short: "A client to interact with the fimserver application",
	}
	showAgentsCmd = &cobra.Command{
		Use:   "show-agents",
		Short: "List all agents registered at the server",
		RunE:  printAgents,
	}
	showAlertsCmd = &cobra.Command{
		Use:   "show-alerts",
		Short: "Print all alerts for a specified agent",
		RunE:  printAlerts,
	}
	approveBaselineUpdateCmd = &cobra.Command{
		Use:   "approve",
		Short: "Approve a baseline update for an agent",
		RunE:  approveBaselineUpdate,
	}
	deleteEndpointCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete the specified endpoint",
		RunE:  deleteEndpoint,
	}
	createAgentCmd = &cobra.Command{
		Use:   "create-agents",
		Short: "Create new agents",
		RunE:  createAgent,
	}
	createClientCmd = &cobra.Command{
		Use:   "create-clients",
		Short: "Create new clients",
		RunE:  createClient,
	}
	updateWatchedPathsCmd = &cobra.Command{
		Use:   "update-paths",
		Short: "Update the watched paths of an agent",
		RunE:  updateWatchedPaths,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&confPath, "config", "config.yaml", "path to config file")

	showAlertsCmd.Flags().StringVar(&agentName, "name", "", "Specify the name of the agent to list alerts for")
	showAlertsCmd.MarkFlagRequired("name")

	approveBaselineUpdateCmd.Flags().StringVar(&agentName, "name", "", "Specify the name of the agent to create an approval for")
	approveBaselineUpdateCmd.MarkFlagRequired("name")

	deleteEndpointCmd.Flags().StringVar(&agentName, "name", "", "Specify the name of the endpoint to delete")
	deleteEndpointCmd.MarkFlagRequired("name")

	createAgentCmd.Flags().StringVar(&filePath, "file", "", "Specify a file containing agent description")
	createAgentCmd.MarkFlagRequired("file")

	createClientCmd.Flags().StringVar(&filePath, "file", "", "Specify a file containing client description")
	createClientCmd.MarkFlagRequired("file")

	updateWatchedPathsCmd.Flags().StringVar(&filePath, "file", "", "Specify a file containing the watched paths and agent name")
	updateWatchedPathsCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(showAgentsCmd, showAlertsCmd)
	rootCmd.AddCommand(approveBaselineUpdateCmd)
	rootCmd.AddCommand(deleteEndpointCmd, createAgentCmd, createClientCmd, updateWatchedPathsCmd)
}
