package command

import (
	"fmt"
	"log"

	"github.com/jrabello/godirzilla/config"
	"github.com/spf13/cobra"
)

var GrpCmd = &cobra.Command{
	Use:   "grp",
	Short: "Manage groups of directories",
}

var grpAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new group",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("You must provide a group name to add")
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		groupName := config.ToGroup(args[0])

		if err := cfg.AddGroup(groupName); err != nil {
			log.Fatalf("Failed to add group: %v", err)
		}
	},
}

var grpRemCmd = &cobra.Command{
	Use:   "rem",
	Short: "Remove a group",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("You must provide a group name to remove")
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		groupName := config.ToGroup(args[0])
		cfg.RemoveGroup(groupName)
	},
}

var grpSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a current group",
	Args:  cobra.ExactArgs(1), // Requires exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		groupName := config.Group(args[0])

		err = cfg.SetCurrentGroup(groupName)
		if err != nil {
			log.Fatalf("Error setting current group: %v", err)
		}
	},
}

var grpListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all groups",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		for groupName := range cfg.Groups {
			fmt.Println(groupName)
		}
	},
}

func init() {
	GrpCmd.AddCommand(grpAddCmd, grpListCmd, grpRemCmd, grpSetCmd)
}
