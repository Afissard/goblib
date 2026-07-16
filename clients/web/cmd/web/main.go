package main

import (
	"os/exec"
	"runtime"
)

func main() {
	url := "http://localhost:8080"

	switch runtime.GOOS {
	case "windows":
		exec.Command("cmd", "/c", "start", url).Start()

	case "darwin":
		exec.Command("open", url).Start()

	default:
		exec.Command("xdg-open", url).Start()
	}
}
