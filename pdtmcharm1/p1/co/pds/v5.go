/*
v5 - This is using codeium directly
*/

package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// main is the entry point of the Go program.
//
// It does not take any parameters.
// It does not return anything.
func V5() {
	// Execute the pdtm command
	output, err := exec.Command("pdtm").CombinedOutput()
	if err != nil {
		log.Fatal("Error executing command:", err)
	}

	// Convert output to string
	outputString := string(output)

	// Parse output and extract tools and statuses
	tools := parseOutput2(outputString)

	// Print tools and statuses with only the first two words
	for tool := range tools {
		// Split the tool into words
		words := strings.Split(tool, " ")
		if len(words) >= 2 {
			toolName := words[0]   // First word
			toolStatus := words[1] // Second word
			fmt.Printf("Tool: %s, Status: %s\n", toolName, toolStatus)
		}
	}
}

// parseOutput parses the output string and extracts tools and statuses.
//
// It takes the output string as a parameter.
// It returns a map of tools and their corresponding statuses.
func parseOutput2(output string) map[string]string {
	lines := strings.Split(output, "\n")
	tools := make(map[string]string)

	// Start reading from the line with tool information
	startReading := false
	for _, line := range lines {
		if strings.HasPrefix(line, "1. ") {
			startReading = true
			continue
		}
		if startReading {
			// Split the line into tool and status
			parts := strings.SplitN(line, " ", 2)
			if len(parts) == 2 {
				tools[parts[1]] = parts[0]
			}
		}
	}

	return tools
}
