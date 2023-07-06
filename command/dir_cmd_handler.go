package command

import (
	"fmt"
	"log"

	"github.com/jrabello/godirzilla/config"
	"github.com/jrabello/godirzilla/file"
	"github.com/spf13/cobra"
)

var DirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Manage directories",
}

var AIDirAddCmd = &cobra.Command{
	Use:   "ai:add [dirPath0] [dirPath1]...",
	Short: "Adds one or more directories to a group that our AI will create for you",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetGlobalConfig()

		if cfg.CurrentGroup == "main" {
			log.Fatal("Group `main` is reserved to run commands only in current directory!!!")
		}

		// Add each directory to the current group
		for _, arg := range args {
			absDir := file.GetAbsoluteFilePath(arg)
			err := cfg.AddDirectoryToGroup(cfg.CurrentGroup, config.ToDirectory(absDir))
			if err != nil {
				log.Fatalf("Error adding directory to group: %v", err)
			}

			fmt.Printf("Added directory `%s` to current group: `%s`\n", absDir, cfg.CurrentGroup)
		}

		err := cfg.ApplyChanges()
		if err != nil {
			log.Fatalf("Error trying to save config file: %v", err)
		}
	},
}

var dirAddCmd = &cobra.Command{
	Use:   "add [dirPath0] [dirPath1]...",
	Short: "Adds one or more directories to the current group",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetGlobalConfig()

		if cfg.CurrentGroup == "" {
			log.Fatal("No current group set")
		}

		if cfg.CurrentGroup == "main" {
			log.Fatal("Group `main` is reserved to run commands only in current directory!!!")
		}

		// Add each directory to the current group
		for _, arg := range args {
			absDir := file.GetAbsoluteFilePath(arg)
			err := cfg.AddDirectoryToGroup(cfg.CurrentGroup, config.ToDirectory(absDir))
			if err != nil {
				log.Fatalf("Error adding directory to group: %v", err)
			}

			fmt.Printf("Added directory `%s` to current group: `%s`\n", absDir, cfg.CurrentGroup)
		}

		err := cfg.ApplyChanges()
		if err != nil {
			log.Fatalf("Error trying to save config file: %v", err)
		}
	},
}

var dirListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all directories in the config file",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetGlobalConfig()
		cfg.PrintDirectoriesFromCurrentGroup()
	},
}

var dirRemCmd = &cobra.Command{
	Use:   "rem [dirName]",
	Short: "Removes a directory from the config file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetGlobalConfig()
		directoryName := config.ToDirectory(args[0])
		cfg.RemoveDirectoryFromCurrentGroup(directoryName)
		cfg.ApplyChanges()
		fmt.Printf("Removed directory `%s` from current group: `%s`\n", directoryName, cfg.CurrentGroup)
	},
}

func init() {
	DirCmd.AddCommand(
		AIDirAddCmd,
		dirAddCmd,
		dirListCmd,
		dirRemCmd,
	)
}
