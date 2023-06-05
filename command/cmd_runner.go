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

	var outbuf, errbuf bytes.Buffer
	c := exec.Command("sh", "-c", fmt.Sprintf("cd %s && %s", dir, command))
	c.Stdout = &outbuf
	c.Stderr = &errbuf
	err := c.Run()

	log.Printf("[thread-%d] Running command '%s' in directory: %s\n", id, command, dir)
	if err != nil {
		errorColor("[thread-%d] Failed to run command in directory %s: %v\n", id, dir, err)
	}
	if outbuf.Len() > 0 {
		log.Printf("[thread-%d] Output:\n%s\n", id, outbuf.String())
	}
	if errbuf.Len() > 0 {
		errorColor("[thread-%d] Errors:\n%s\n", id, errbuf.String())
	}
}

// Run all commands in their respective target directories
func CreateThreadsAndRunAllCommands(command string, directories []config.Directory) {
	var wg sync.WaitGroup
	goroutineID := new(int32)
	for _, dir := range directories {
		wg.Add(1)
		go runCommand(dir, command, int(atomic.AddInt32(goroutineID, 1)-1), &wg)
	}
	// wait for all goroutines to finish
	wg.Wait()
}
