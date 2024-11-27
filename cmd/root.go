/*
Copyright 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "taskmanager <command>",
	Short: "A simple task manager written in Go.",
	Long: `This task manager is a simple CLI application that allows you 
	 to manage your tasks. It is written in Go and uses Cobra to build a 
	 command line interface. It also uses a JSON file to store the tasks.`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Check if the user provided a command
		if len(args) == 0 {
			return fmt.Errorf("Please provide a command. Use 'taskmanager help' for more information.")
		}
		return nil
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.taskmanager.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
