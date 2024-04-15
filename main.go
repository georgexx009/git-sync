package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
  config, newContent := readConfigFile()
  if !config.on {
    return
  }
	executeCommand("git", "add", ".")
	executeCommand("git", "commit", "-m", "sync")
	executeCommand("git", "push", "origin", "main")
  config.count++
  if config.count >= config.timeToReset {
    config.count = 0
  }
	fmt.Println("Synced")
  newContent = strings.Replace(newContent, "COUNT=" + strconv.Itoa(config.count - 1), "COUNT=" + strconv.Itoa(config.count), 1)
  updateFile(newContent)
}

func executeCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command finished with error: %v", err)
	}
}

type Config struct {
	on          bool
	timeToReset int
	count       int
}

type Validator struct {
  on bool
  timeToReset bool
  count bool
}

func readConfigFile() (Config, string){
	file, err := os.Open("./.git-sync")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()
	config := Config{}
  newContent := ""

	scanner := bufio.NewScanner(file)
  validator := Validator{on: false, timeToReset: false, count: false}
	for scanner.Scan() {
		str := scanner.Text()
		elements := strings.Split(str, "=")
		base := elements[0]
		value := elements[1]
		switch base {
		case "ON":
      validator.on = true
			config.on = value == "true"
      newContent += value + "\n"
		case "TIME_TO_RESET":
      valInt, err := strconv.Atoi(value)
      if err != nil {
        log.Fatalf("Error converting string to int: %v", err)
      }
      validator.timeToReset = true
			config.timeToReset = valInt
      newContent += value + "\n"
		case "COUNT":
      valInt, err := strconv.Atoi(value)
      if err != nil {
        log.Fatalf("Error converting string to int: %v", err)
      }
      validator.count = true
			config.count = valInt
      newContent += value + "\n"
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

  if !validator.on {
    log.Fatalf("ON field is required")
  }
  if !validator.timeToReset {
    log.Fatalf("TIME_TO_RESET field is required")
  }
  if !validator.count {
    log.Fatalf("COUNT field is required")
  }

  return config, newContent
}

func updateFile(newContent string) {
  err := os.WriteFile("./.git-sync", []byte(newContent), 0644)
  if err != nil {
    log.Fatalf("Error writing file: %v", err)
  }
}
