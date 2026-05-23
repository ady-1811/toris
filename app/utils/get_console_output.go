package utils

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

func GetConsoleOutput() (string, error) {
	os := GetOS()
	switch {
	case strings.Contains(os, "linux"):
		return getLinuxTerminalBuffer()
	case strings.Contains(os, "darwin"): // macOS and Linux
		return getmacOSTerminalBuffer()
	case strings.Contains(os, "windows"):
		return getWindowsTerminalBuffer()
	default:
		return "", fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

// AI-generated code
func getmacOSTerminalBuffer() (string, error) {
	script := `tell application "Terminal" to get contents of selected tab of window 1`
	cmd := exec.Command("osascript", "-e", script)
	out, err := cmd.Output()
	return cleanTerminalText(string(out)), err
}

// AI-generated code
func getLinuxTerminalBuffer() (string, error) {
	logFile := "/tmp/toris.log"
	if envLog := os.Getenv("TORIS_LOG"); envLog != "" {
		logFile = envLog
	}

	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		return "", fmt.Errorf("linux capture requires a recorded session.\n\nTo automate this setup, please run:\n  toris init\n\nOr to start manually for this session, type:\n  script %s", logFile)
	}

	// Read a larger chunk to ensure we get the full last command output
	cmd := exec.Command("tail", "-n", "1000", logFile)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to read script log: %v", err)
	}

	logs := string(out)
	parts := strings.Split(logs, "---TORIS_CMD_START---")
	if len(parts) >= 2 {
		// The very last part is the currently executing 'toris scan'. We want the block right before it!
		logs = parts[len(parts)-2]
	}
	return cleanTerminalText(logs), nil
}

// AI-generated code
func getWindowsTerminalBuffer() (string, error) {
	psCommand := `
		$Host.UI.RawUI.GetBufferContents(
			(New-Object System.Management.Automation.Host.Rectangle(0,0,80,25))
		) | ForEach-Object { $_.Character } -join ''
	`
	cmd := exec.Command("powershell", "-Command", psCommand)
	out, err := cmd.Output()
	return cleanTerminalText(string(out)), err
}

// AI-generated code
func cleanTerminalText(input string) string {
	// Matches CSI escapes (colors) and OSC sequences (terminal titles, shell integrations)
	re := regexp.MustCompile(`\x1b\[[0-9;?]*[a-zA-Z]|\x1b\][^\x07\x1b]*(?:\x07|\x1b\\)`)
	clean := re.ReplaceAllString(input, "")

	return strings.TrimSpace(strings.ReplaceAll(clean, "\r", ""))
}
