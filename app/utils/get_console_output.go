package utils

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"regexp"
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
	cmd := exec.Command("tmux", "capture-pane", "-p")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("linux capture requires running inside TMUX")
	}
	return cleanTerminalText(string(out)), err
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
	const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[-a-zA-Z\\d\\/#&.:=?%@~]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PR-TZcf-ntqry=><~]))"
	var re = regexp.MustCompile(ansi)
	clean := re.ReplaceAllString(input, "")
	
	return strings.TrimSpace(clean)
}

