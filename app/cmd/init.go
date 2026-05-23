/*
Copyright © 2025 TORIS
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/toris/utils"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize TORIS auto-recording in your shell",
	Long: `Injects a small script into your ~/.bashrc and/or ~/.zshrc file
so that terminal sessions are automatically recorded for the 'toris scan' feature.`,
	Run: func(cmd *cobra.Command, args []string) {
		injected, err := utils.SetupShellRecording()
		if err != nil {
			fmt.Println(err)
			return
		}

		if !injected {
			fmt.Println("No changes made. Make sure you have a ~/.bashrc or ~/.zshrc file, or the snippet might already exist.")
		} else {
			fmt.Println("\nPlease restart your terminal or run 'source ~/.bashrc' (or ~/.zshrc) for changes to take effect.")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
