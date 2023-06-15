package command

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/jrabello/godirzilla/config"
	"github.com/spf13/cobra"
)

var DirCmd = &cobra.Command{
	Use:   "dir",
	Short: "Manage directories",
}

var dirAddCmd = &cobra.Command{
	Use:   "add [dir]",
	Short: "Adds a directory to the current group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		partialDir := args[0]

		// Convert the partial directory to absolute directory
		absDir, err := filepath.Abs(partialDir)
		if err != nil {
			log.Fatalf("Error converting partial directory to absolute: %v", err)
		}

		// Check if the absolute directory exists
		_, err = os.Stat(absDir)
		if os.IsNotExist(err) {
			log.Fatalf("Directory %s does not exist", absDir)
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		// Check if there is a current group set
		if cfg.CurrentGroup == "" {
			log.Fatal("No current group set")
		}

		// Add the directory to the current group
		err = cfg.AddDirectoryToGroup(cfg.CurrentGroup, config.ToDirectory(absDir))
		if err != nil {
			log.Fatalf("Error adding directory to group: %v", err)
		}

		err = cfg.ApplyChanges()
		if err != nil {
			log.Fatalf("Error trying to save config file: %v", err)
		}

		fmt.Printf("Added directory %s to current group\n", absDir)
	},
}

// var dirListCmd = &cobra.Command{
// 	Use:   "list",
// 	Short: "Lists all directories in the config file",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		HandleListCommand()
// 	},
// }

// var dirRemCmd = &cobra.Command{
// 	Use:   "rem [dir]",
// 	Short: "Removes a directory from the config file",
// 	Args:  cobra.ExactArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		HandleRemoveCommand(args)
// 	},
// }

func init() {
	DirCmd.AddCommand(
		dirAddCmd,
		// dirListCmd,
		// dirRemCmd
	)
}
