package command

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login via Firebase OAuth",
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://godirzilla.cognuscraft.com/login"

		// Open the URL in the user's browser
		openBrowser(url)

		// Start a server to listen on a local port
		go startServer()
	},
}

func openBrowser(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	if err := cmd.Start(); err != nil {
		fmt.Printf("Failed to open browser, visit the following URL yourself: %v", url)
	}
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// this will be hit when your login service redirects the user back to localhost
		fmt.Fprint(w, "Login successful!")
		// At this point you can close the server and store the user credentials securely
	})
	http.ListenAndServe(":8080", nil)
}
