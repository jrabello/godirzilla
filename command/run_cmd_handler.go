package command

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jrabello/godirzilla/config"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run <command> [args...]",
	Short: "Runs a command in all directories in the config file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetGlobalConfig()

		if cfg.CurrentGroup == "" {
			log.Fatal("No current group is set in config, please set a current group before running a command!")
		}

		command := strings.Join(args, " ")
		directoryList, err := cfg.GetDirectoriesFromGroup(cfg.CurrentGroup)
		if err != nil {
			log.Fatalf("No DirectoryList was found for group: %s", cfg.CurrentGroup)
		}

		if cfg.CurrentGroup == "main" {
			currentDir, err := os.Getwd()
			if err != nil {
				log.Fatalf("Unable to get current directory!")
			}

			directoryList = append(directoryList, config.ToDirectory(currentDir))
		}

		if len(directoryList) == 0 {
			fmt.Printf("I can not run the command: `%s` inside group: `%s` because it has no directories, please add at least one directory in order to run commands into it!\n",
				command, cfg.CurrentGroup)
			return
		}

		CreateThreadsAndRunAllCommands(command, cfg.CurrentGroup, directoryList)
	},
}
