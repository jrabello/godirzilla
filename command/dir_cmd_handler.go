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

var dirAddCmd = &cobra.Command{
	Use:   "add [dirName]",
	Short: "Adds a directory to the current group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		// Check if there is a current group set
		if cfg.CurrentGroup == "" {
			log.Fatal("No current group set")
		}

		// Add the directory to the current group
		absDir := file.GetAbsoluteFilePath(args[0])
		err = cfg.AddDirectoryToGroup(cfg.CurrentGroup, config.ToDirectory(absDir))
		if err != nil {
			log.Fatalf("Error adding directory to group: %v", err)
		}

		err = cfg.ApplyChanges()
		if err != nil {
			log.Fatalf("Error trying to save config file: %v", err)
		}
		fmt.Printf("Added directory %s to current group: %s\n", absDir, cfg.CurrentGroup)
	},
}

var dirListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all directories in the config file",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}
		cfg.PrintDirectoriesFromCurrentGroup()
	},
}

var dirRemCmd = &cobra.Command{
	Use:   "rem [dirName]",
	Short: "Removes a directory from the config file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}
		directoryName := config.ToDirectory(args[0])
		cfg.RemoveDirectoryFromCurrentGroup(directoryName)
		cfg.ApplyChanges()
		fmt.Printf("Removed directory `%s` from current group: `%s`\n", directoryName, cfg.CurrentGroup)
	},
}

func init() {
	DirCmd.AddCommand(
		dirAddCmd,
		dirListCmd,
		dirRemCmd,
	)
}
