package command

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/fatih/color"
	"github.com/jrabello/godirzilla/config"
	"github.com/jrabello/godirzilla/utils"
)

type CommandResult struct {
	DirPath    string
	ExitStatus int
	isError    bool
	Stdout     string
	Stderr     string
}

func runCommand(dirPath config.Directory, command string, threadId int) CommandResult {
	log.SetFlags(0)

	c := exec.Command("sh", "-c", fmt.Sprintf("cd %s && %s", dirPath, command))
	var stdoutbuf, stderrbuf bytes.Buffer
	c.Stdout = &stdoutbuf
	c.Stderr = &stderrbuf
	err := c.Run()

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	dirBase := filepath.Base(string(dirPath))

	var exitCode int = 0
	unknownError := false
	errorDescription := ""
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			// The process returned a non-zero exit code
			exitCode = exitError.ExitCode()
			errorDescription = fmt.Sprintf("error:\nfailed to run command in directory %s. Exit code: %d\n\n", dirPath, exitCode)
		} else {
			// The command execution encountered an error
			unknownError = true
			errorDescription = fmt.Sprintf("error:\nfailed to run command in directory %s. Error: %v %t\n\n", dirPath, err, unknownError)
		}
	}

	rightPadding := utils.PadRight(dirBase, ".", 25)
	if exitCode == 0 {
		fmt.Printf("📂 %s%s\n", rightPadding, green("✔️"))
	} else {
		fmt.Printf("📂 %s%s (Exit %d)\n\n", rightPadding, red("❌"), exitCode)
	}

	if stdoutbuf.Len() > 0 {
		fmt.Printf("%s\n", stdoutbuf.String())
	}
	if stderrbuf.Len() > 0 {
		fmt.Printf("%s\n", stderrbuf.String())
	}
	if len(errorDescription) > 0 {
		fmt.Println(errorDescription)
	}

	return CommandResult{
		DirPath:    string(dirPath),
		ExitStatus: exitCode,
		isError:    (unknownError || exitCode > 0),
		Stdout:     stdoutbuf.String(),
		Stderr:     stderrbuf.String(),
	}
}

// CreateThreadsAndRunAllCommands runs all commands in their respective group and target directories in separate threads
func CreateThreadsAndRunAllCommands(command string, currentGroup config.Group, directories []config.Directory) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("🚀 Running command: '%s' \n", red(command))
	fmt.Printf("   Group: '%s' \n\n", red(currentGroup))

	var wg sync.WaitGroup
	goroutineID := new(int32)

	// Create a buffered channel with capacity equal to number of directories
	results := make(chan CommandResult, len(directories))

	for _, dir := range directories {
		wg.Add(1)
		go func(dir config.Directory, results chan<- CommandResult) {
			defer wg.Done()
			results <- runCommand(dir, command, int(atomic.AddInt32(goroutineID, 1)-1))
		}(dir, results)
	}

	// Wait for all goroutines to finish and close the results channel
	wg.Wait()
	close(results)

	// Read from the results channel until it's empty
	successCount := 0
	failureCount := 0
	for result := range results {
		if result.ExitStatus == 0 {
			successCount++
		} else {
			failureCount++
		}
	}

	fmt.Println("\n📊 Summary:")
	fmt.Printf("Successful Operations: %d\n", successCount)
	fmt.Printf("Failed Operations: %d\n\n", failureCount)
}
