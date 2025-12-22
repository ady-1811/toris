package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetConsoleOutput() (string, error) {
	os := GetOS()

	switch {
	case strings.Contains(os, "linux"), strings.Contains(os, "darwin"): // macOS and Linux
		return getUnixHistory()
	case strings.Contains(os, "windows"):
		return getWindowsPowerShellHistory()
	default:
		return "", fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

// AI-Generated
func getUnixHistory() (string, error) {
	home, _ := os.UserHomeDir()
	shell := os.Getenv("SHELL")
	var historyFile string

	if strings.Contains(shell, "zsh") {
		historyFile = home + "/.zsh_history"
	} else {
		historyFile = home + "/.bash_history"
	}

	file, err := os.Open(historyFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	stat, _ := file.Stat()
	start := stat.Size() - 500
	if start < 0 {
		start = 0
	}

	buf := make([]byte, 500)
	_, err = file.ReadAt(buf, start)
	return string(buf), err
}

// AI-Generated
func getWindowsPowerShellHistory() (string, error) {
	cmd := exec.Command("powershell", "-Command", "Get-History -Count 5 | Select-Object -ExpandProperty CommandLine")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
