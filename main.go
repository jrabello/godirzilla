package main

import (
	"fmt"
	"os"

	"github.com/jrabello/godirzilla/command"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godirzilla",
	Short: "GoDirZilla is a powerful CLI tool to execute the same command in multiple directories.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	var addCmd = &cobra.Command{
		Use:   "add [dir]",
		Short: "Adds a directory to the config file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			command.HandleAddCommand(args)
		},
	}
	rootCmd.AddCommand(addCmd)

	var remCmd = &cobra.Command{
		Use:   "rem [dir]",
		Short: "Removes a directory from the config file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			command.HandleRemoveCommand(args)
		},
	}
	rootCmd.AddCommand(remCmd)

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all directories in the config file",
		Run: func(cmd *cobra.Command, args []string) {
			command.HandleListCommand()
		},
	}
	rootCmd.AddCommand(listCmd)

	var runCmd = &cobra.Command{
		Use:   "run <command> [args...]",
		Short: "Runs a command in all directories in the config file",
		Run: func(cmd *cobra.Command, args []string) {
			command.HandleRunCommand(args)
		},
	}
	rootCmd.AddCommand(runCmd)
	Execute()
}
