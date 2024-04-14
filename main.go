package main

import (
	"log"
	"os/exec"
)

func executeCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command finished with error: %v", err)
	}
}

func main() {
	executeCommand("git", "add", ".")
	executeCommand("git", "commit", "-m", "sync")
	executeCommand("git", "push", "origin", "main")
}

