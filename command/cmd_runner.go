package command

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"sync"
	"sync/atomic"

	"github.com/fatih/color"
	"github.com/jrabello/godirzilla/config"
)

func runCommand(dir config.Directory, command string, id int, wg *sync.WaitGroup) {
	log.SetFlags(0)
	defer wg.Done()

	errorColor := color.New(color.FgRed).PrintfFunc()
	infoColor := color.New(color.FgWhite).PrintfFunc()

	log.Printf("[thread-%d] Running command: '%s' in directory: '%s'\n", id, command, dir)
	c := exec.Command("sh", "-c", fmt.Sprintf("cd %s && %s", dir, command))
	var outbuf, errbuf bytes.Buffer
	c.Stdout = &outbuf
	c.Stderr = &errbuf
	err := c.Run()

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			// The process returned a non-zero exit code
			exitCode := exitError.ExitCode()
			errorColor("[thread-%d] Failed to run command in directory %s. Exit code: %d\n", id, dir, exitCode)
		} else {
			// The command execution encountered an error
			errorColor("[thread-%d] Failed to run command in directory %s. Error: %v\n", id, dir, err)
		}
	}

	if outbuf.Len() > 0 {
		infoColor("[thread-%d] stdout:\n%s\n", id, outbuf.String())
	}

	if errbuf.Len() > 0 {
		errorColor("[thread-%d] stderr:\n%s\n", id, errbuf.String())
	}
}

// CreateThreadsAndRunAllCommands runs all commands in their respective group and target directories in separate threads
func CreateThreadsAndRunAllCommands(command string, directories []config.Directory) {
	var wg sync.WaitGroup
	goroutineID := new(int32)
	for _, dir := range directories {
		wg.Add(1)
		go runCommand(dir, command, int(atomic.AddInt32(goroutineID, 1)-1), &wg)
	}
	wg.Wait()
}
