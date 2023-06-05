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

func HandleAddCommand(args []string) {
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

	// Add the directory to the configuration
	cfg.Directories = append(cfg.Directories, absDir)

	// Save the updated configuration
	err = config.SaveConfig(cfg)
	if err != nil {
		log.Fatalf("Error saving config: %v", err)
	}

	fmt.Printf("Added directory %s to config\n", absDir)
}

func HandleRemoveCommand(args []string) {
	dir := args[0]
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Find and remove the directory from the configuration
	for i, d := range cfg.Directories {
		if d == dir {
			cfg.Directories = append(cfg.Directories[:i], cfg.Directories[i+1:]...)
			break
		}
	}

	// Save the updated configuration
	err = config.SaveConfig(cfg)
	if err != nil {
		log.Fatalf("Error saving config: %v", err)
	}

	fmt.Printf("Removed directory %s from config\n", dir)
}

func HandleListCommand() {
	fmt.Println("Listing directories from config file...")

	// Load the configuration
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	// Print the directories
	fmt.Printf("Found %d directories!\n", len(config.Directories))
	for _, dir := range config.Directories {
		fmt.Println(dir)
	}
}

var dirAddCmd = &cobra.Command{
	Use:   "add [dir]",
	Short: "Adds a directory to the config file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		HandleAddCommand(args)
	},
}

var dirListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all directories in the config file",
	Run: func(cmd *cobra.Command, args []string) {
		HandleListCommand()
	},
}

var dirRemCmd = &cobra.Command{
	Use:   "rem [dir]",
	Short: "Removes a directory from the config file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		HandleRemoveCommand(args)
	},
}

func init() {
	DirCmd.AddCommand(dirAddCmd, dirListCmd, dirRemCmd)
}
