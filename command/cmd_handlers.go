package command

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"

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

func runCommand(dir, command string, id int, wg *sync.WaitGroup) {
	defer wg.Done() // make sure this is the first line in your goroutine, so it gets called when the function exits

	shell := os.Getenv("SHELL")
	var sourceFile string
	switch {
	case strings.Contains(shell, "bash"):
		sourceFile = "~/.bashrc"
	case strings.Contains(shell, "zsh"):
		sourceFile = "~/.zshrc"
	default:
		// Default to bash if the shell isn't recognized
		sourceFile = "~/.bashrc"
	}

	var outbuf, errbuf bytes.Buffer
	c := exec.Command(shell, "-c", fmt.Sprintf("source %s; cd %s && %s", sourceFile, dir, command))
	c.Stdout = &outbuf
	c.Stderr = &errbuf
	err := c.Run()

	log.Printf("[Thread %d] Running command '%s' in directory: %s\n", id, command, dir)
	if err != nil {
		log.Printf("[Thread %d] Failed to run command in directory %s: %v\n", id, dir, err)
	}
	if outbuf.Len() > 0 {
		log.Printf("[Thread %d] Output:\n%s\n", id, outbuf.String())
	}
	if errbuf.Len() > 0 {
		log.Printf("[Thread %d] Errors:\n%s\n", id, errbuf.String())
	}
}

func HandleRunCommand(args []string) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	command := strings.Join(args, " ")

	var wg sync.WaitGroup
	goroutineID := new(int32)
	for _, dir := range cfg.Directories {
		wg.Add(1)
		go runCommand(dir, command, int(atomic.AddInt32(goroutineID, 1)-1), &wg) // pass the wait group to your function
	}
	wg.Wait() // wait for all goroutines to finish
}
