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
	rootCmd.AddCommand(command.DirCmd)
	rootCmd.AddCommand(command.GrpCmd)
	rootCmd.AddCommand(command.RunCmd)
	Execute()
}
