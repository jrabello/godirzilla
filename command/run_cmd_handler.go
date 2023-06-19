package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/jrabello/godirzilla/config"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run <command> [args...]",
	Short: "Runs a command in all directories in the config file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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

		fmt.Printf("Running command: `%s` inside directories of group: `%s`\n", command, cfg.CurrentGroup)
		CreateThreadsAndRunAllCommands(command, directoryList)
	},
}
