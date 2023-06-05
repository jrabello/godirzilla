package command

import (
	"log"
	"strings"

	"github.com/jrabello/godirzilla/config"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run <command> [args...]",
	Short: "Runs a command in all directories in the config file",
	Run: func(cmd *cobra.Command, args []string) {
		HandleRunCommand(args)
	},
}

func HandleRunCommand(args []string) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	command := strings.Join(args, " ")

	CreateThreadsAndRunAllCommands(command, cfg.Directories)
}
