package command

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/jrabello/godirzilla/config"
)

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

func HandleRunCommand(args []string) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	command := strings.Join(args, " ")

	var wg sync.WaitGroup
	for _, dir := range config.Directories {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			fmt.Printf("Running command in directory: %s\n", dir)
			c := exec.Command("sh", "-c", fmt.Sprintf("cd %s && %s", dir, command))
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
			if err := c.Run(); err != nil {
				log.Printf("Failed to run command in directory %s: %v", dir, err)
			}
		}(dir)
	}
	wg.Wait()
}
