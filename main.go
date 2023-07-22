package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jrabello/godirzilla/command"
	"github.com/jrabello/godirzilla/config"
	"github.com/spf13/cobra"
)

type MyApp struct {
	RootCmd *cobra.Command
	Config  *config.Config
}

func (app *MyApp) ExecRootCmd() {
	if err := app.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initializeApplication(cfg *config.Config) *MyApp {
	app := &MyApp{
		Config: cfg,
		RootCmd: &cobra.Command{
			Use:   "gdz",
			Short: "GoDirZilla is a fast, cool and cognitive CLI tool to execute commands in groups of directories at the same time with a rich output that makes you really understand what's going on in your terminal",
		},
	}

	app.RootCmd.AddCommand(command.GetDirCmd(app.Config))
	app.RootCmd.AddCommand(command.GrpCmd(app.Config))
	app.RootCmd.AddCommand(command.RunCmd(app.Config))
	app.RootCmd.AddCommand(command.AICmd(app.Config))

	return app
}

func main() {
	config, err := config.LoadGlobalConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	app := initializeApplication(config)
	app.ExecRootCmd()
}
