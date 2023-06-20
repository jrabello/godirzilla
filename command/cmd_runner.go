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
)

func PadRight(str, pad string, length int) string {
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}

func runCommand(dirPath config.Directory, command string, threadId int, wg *sync.WaitGroup) {
	log.SetFlags(0)
	defer wg.Done()

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
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			// The process returned a non-zero exit code
			exitCode = exitError.ExitCode()
			fmt.Printf("error:\nfailed to run command in directory %s. Exit code: %d\n\n", dirPath, exitCode)
		} else {
			// The command execution encountered an error
			unknownError = true
			fmt.Printf("error:\nfailed to run command in directory %s. Error: %v %d\n\n", dirPath, err, unknownError)
		}
	}

	if exitCode == 0 {
		fmt.Printf("📂 %s%s\n\n", PadRight(dirBase, ".", 25), green("✔️"))
	} else {
		fmt.Printf("📂 %s%s (Exit %d)\n\n", PadRight(dirBase, ".", 25), red("❌"), exitCode)
	}

	if stdoutbuf.Len() > 0 {
		// fmt.Printf("📂 %s:\n%s\n", dirBase, stdoutbuf.String())
		fmt.Printf("%s\n", stdoutbuf.String())
	}
	if stderrbuf.Len() > 0 {
		fmt.Printf("%s\n", stderrbuf.String())
		// fmt.Printf("📂 %s%s\n", PadRight(dirBase, ".", 25), red("❌"))
		// fmt.Printf("stderr:\n%s\n", stderrbuf.String())
	}
}

// CreateThreadsAndRunAllCommands runs all commands in their respective group and target directories in separate threads
func CreateThreadsAndRunAllCommands(command string, currentGroup config.Group, directories []config.Directory) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("🚀 Running command: '%s' in group: '%s'...\n\n", red(command), red(currentGroup))

	var wg sync.WaitGroup
	goroutineID := new(int32)
	for _, dir := range directories {
		wg.Add(1)
		go runCommand(dir, command, int(atomic.AddInt32(goroutineID, 1)-1), &wg)
	}

	wg.Wait()
}
