package utils

import (
	"fmt"
)

func PrintInfo(osName string, commands []string, confidence float64, instructions []string, risk_score int, confirmation_required bool) {
	fmt.Printf("OS: %s\n", osName)

	fmt.Printf("Commands: \n")
	for i := 1; i <= len(commands); i++ {
		fmt.Printf("%v: %s\n", i, commands[i-1])
	}

	fmt.Printf("Confidence: %.2f\n", confidence)

	fmt.Printf("Instructions: \n")
	for i := 1; i <= len(instructions); i++ {
		fmt.Printf("%v: %s\n", i, instructions[i-1])
	}

	fmt.Printf("Confidence: %v\n", risk_score)
	fmt.Printf("Confirmation Required: %v\n", confirmation_required)
}
