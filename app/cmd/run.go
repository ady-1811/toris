/*
Copyright © 2025 TORIS

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/toris/ai"
	"github.com/toris/utils"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a command based on your description",
	Long: `Run translates your natural language description into a terminal command
and executes it in your shell environment. For example:

  toris run "List all files in the current directory"

This will generate and execute the appropriate command to list files.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client, err := ai.NewGeminiCommandClient("gemini-2.5-flash")
		if err != nil {
			log.Fatalf("Init error: %v", err)
		}

		result, err := client.GetCommand(ctx, args[0])
		if err != nil {
			log.Fatalf("API error: %v", err)
		}

		utils.PrintInfo(client.OSName, result.Command, result.Confidence, result.Instruction)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
