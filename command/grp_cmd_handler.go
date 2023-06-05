package command

import (
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
		// Add your logic here
	},
}

var grpListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all groups",
	Run: func(cmd *cobra.Command, args []string) {
		// Add your logic here
	},
}

var grpRemCmd = &cobra.Command{
	Use:   "rem",
	Short: "Remove a group",
	Run: func(cmd *cobra.Command, args []string) {
		// Add your logic here
	},
}

var grpSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a group for a directory",
	Run: func(cmd *cobra.Command, args []string) {
		// Add your logic here
	},
}

func init() {
	GrpCmd.AddCommand(grpAddCmd, grpListCmd, grpRemCmd, grpSetCmd)
}
