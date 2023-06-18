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
	Use:   "add [groupName]",
	Short: "Add a new group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		groupName := config.ToGroup(args[0])
		if err := cfg.AddGroup(groupName); err != nil {
			log.Fatalf("Failed to add group: %v", err)
		}

		cfg.ApplyChanges()
		fmt.Printf("group: `%s` created successfully!\n", groupName)
	},
}

var grpRemCmd = &cobra.Command{
	Use:   "rem [groupName]",
	Short: "Remove a group",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			log.Fatalf("Error loading config: %v", err)
		}

		groupName := config.ToGroup(args[0])
		cfg.RemoveGroup(groupName)
		cfg.ApplyChanges()
		fmt.Printf("group: `%s` removed successfully!\n", groupName)
	},
}

var grpSetCmd = &cobra.Command{
	Use:   "set [groupName]",
	Short: "Set a current group to work with",
	Args:  cobra.ExactArgs(1),
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
		cfg.ApplyChanges()
		fmt.Printf("group: `%s` set as current successfully!\n", groupName)
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

		if len(cfg.Groups) == 0 {
			fmt.Println("No group found!")
		}

		fmt.Println("groups:")
		for groupName := range cfg.Groups {
			if groupName == cfg.CurrentGroup {
				fmt.Printf("* %s (current)\n", groupName)
			} else {
				fmt.Println(groupName)
			}
		}
	},
}

func init() {
	GrpCmd.AddCommand(grpAddCmd, grpListCmd, grpRemCmd, grpSetCmd)
}
