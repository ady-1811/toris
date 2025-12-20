package utils

import (
	"fmt"
)

func PrintInfo(osName string, command string, confidence float64, instructions []string) {
	fmt.Printf("OS: %s\n", osName)
	fmt.Printf("Suggested Command: %s\n", command)
	fmt.Printf("Confidence: %.2f\n", confidence)

	fmt.Printf("Instructions: \n")
	for i := 1; i < len(instructions); i++ {
		fmt.Printf("%v: %s\n", i, instructions[i])
	}
}
