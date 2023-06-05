package main

import (
	"fmt"
	"os"

	"github.com/jrabello/godirzilla/command"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "godirzilla",
	Short: "GoDirZilla is a cool and fast CLI tool to execute the same command in groups of directories at the same time in separate threads",
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
