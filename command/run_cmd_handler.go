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

	if cfg.CurrentGroup == "" {
		log.Fatal("No current group is set in config, please set a current group before running a command!")
	}

	command := strings.Join(args, " ")
	directoryList, err := cfg.GetDirectoriesFromGroup(cfg.CurrentGroup)
	if err != nil {
		log.Fatalf("No DirectoryList was found for group: %s", cfg.CurrentGroup)
	}

	CreateThreadsAndRunAllCommands(command, directoryList)
}
