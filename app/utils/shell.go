package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const InitSnippet = `
# TORIS Auto-recording
if [ -z "$TORIS_RECORDING" ]; then
    export TORIS_RECORDING=1
    script -q -a /tmp/toris.log
    exit
fi

# TORIS latest command marker
if [ -n "$BASH_VERSION" ]; then
    export PS0="---TORIS_CMD_START---\n"
elif [ -n "$ZSH_VERSION" ]; then
    autoload -Uz add-zsh-hook 2>/dev/null
    _toris_marker() { printf "---TORIS_CMD_START---\n"; }
    add-zsh-hook preexec _toris_marker 2>/dev/null
fi
`

func SetupShellRecording() (bool, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return false, fmt.Errorf("error finding home directory: %v", err)
	}

	shells := []string{".bashrc", ".zshrc"}
	injected := false
	for _, shell := range shells {
		rcPath := filepath.Join(home, shell)
		if _, err := os.Stat(rcPath); err == nil {
			if injectSnippet(rcPath) {
				injected = true
			}
		}
	}
	return injected, nil
}

func injectSnippet(rcPath string) bool {
	content, err := os.ReadFile(rcPath)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", rcPath, err)
		return false
	}

	if strings.Contains(string(content), "TORIS_RECORDING") {
		fmt.Printf("Snippet already exists in %s\n", rcPath)
		return false
	}

	f, err := os.OpenFile(rcPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", rcPath, err)
		return false
	}
	defer f.Close()

	if _, err := f.WriteString(InitSnippet); err != nil {
		fmt.Printf("Error writing to %s: %v\n", rcPath, err)
		return false
	}

	fmt.Printf("Successfully added TORIS snippet to %s\n", rcPath)
	return true
}
