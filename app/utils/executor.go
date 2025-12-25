package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

)

type Executor struct {
	Commands [][]string
	Confirm  bool
}

var lastOutput string

func NewExecutor(commandResponse CommandResponse) *Executor {
	var commands [][]string
	for _, cmdStr := range commandResponse.Command {
		parts := strings.Fields(cmdStr)
		if len(parts) > 0 {
			commands = append(commands, parts)
		}
	}
	return &Executor{
		Commands: commands,
		Confirm:  commandResponse.Confirm,
	}
}

func (e *Executor) ExecuteAll() error {
	if e.Confirm {
		GetUserConfirmation("One of these commands may be critical and need your attention, please confrim to proceed: ")
	}

	if !e.Confirm {
		for _, cmdParts := range e.Commands {
			if len(cmdParts) == 0 {
				continue
			}
			cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
			output, err := cmd.CombinedOutput()
			if err != nil {
				return err
			}
			lastOutput = string(output)
			println(string(output))
		}
		return nil
	} else {
		return fmt.Errorf("need user confirmation to execute critical commands!")
	}
}

func GetLastOutput() (string, error) {
	if lastOutput == "" {
		return "", fmt.Errorf("no command has been executed yet")
	}
	return lastOutput, nil
}

func GetUserConfirmation(prompt string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [y/n]: ", prompt)
		response, err := reader.ReadString('\n')
		if err != nil {
			return false
		}

		response = strings.ToLower(strings.TrimSpace(response))
		switch response {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		}

		fmt.Println("Please type 'y' or 'n'.")
	}
}

func ExecuteCommand(result *CommandResponse) {
	executor := NewExecutor(*result)
	err := executor.ExecuteAll()
	if err != nil {
		log.Fatalf("Execution error: %v", err)
	}
}
