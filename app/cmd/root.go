/*
Copyright © 2025 TORIS
*/
package cmd

import (
	"os"
	"log"

	"github.com/spf13/cobra"
	"github.com/toris/ai"
)

var Client *ai.GeminiCommandClient

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hello",
	Short: "TORIS - Terminal Organized and Rational IntelliSense",
	Long: `TORIS is an AI-powered terminal assistant designed to enhance your command-line experience.
With TORIS, you can get intelligent suggestions, automate tasks, and streamline your workflow 
directly from the terminal.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		Client, err = ai.NewGeminiCommandClient("gemini-2.5-flash")
		if err != nil {
			log.Fatalf("Init error: %v", err)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hello.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
