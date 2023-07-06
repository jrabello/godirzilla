package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jrabello/godirzilla/command"
	"github.com/jrabello/godirzilla/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gdz",
	Short: "GoDirZilla is a cool and fast CLI tool to execute the same command in groups of directories at the same time with a rich output that makes you really understand whats going on",
}

func execRootCmd() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	err := config.LoadGlobalConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	rootCmd.AddCommand(command.DirCmd)
	rootCmd.AddCommand(command.GrpCmd)
	rootCmd.AddCommand(command.RunCmd)

	execRootCmd()
}
