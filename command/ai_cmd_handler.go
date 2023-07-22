package command

import (
	"github.com/jrabello/godirzilla/config"
	"github.com/spf13/cobra"
)

var AICmd = &cobra.Command{
	Use:   "ai",
	Short: "Uses artificial inteligence to manage your group names for you, so you don't have to think about it!",
}

var aiDirAddCmd = &cobra.Command{
	Use:   "add [dirPath0] [dirPath1]...",
	Short: "Adds one or more directories to a new group that our AI will create for you",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetGlobalConfig()

		cfg.ApplyChanges()
	},
}

func init() {
	AICmd.AddCommand(dirAddCmd)
}
