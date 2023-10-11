/*
This is a continuation of methods.go to prevent weird errors
- This is from GPT3 , and its much better
*/

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

// // =========================================================
// //version 1

// // Output struct
// type Output struct {
// 	Tools map[string]string
// }

// func Met3() {
// 	// Execute PDTM command
// 	out, err := exec.Command("pdtm").Output()
// 	if err != nil {
// 		fmt.Println("Error executing command:", err)
// 		return
// 	}

// 	// Convert output to string
// 	output := string(out)

// 	// Parse output and extract tools and statuses
// 	tools := parseOutput(output)

// 	// Print tools and statuses
// 	for tool, status := range tools {
// 		fmt.Printf("%s %s\n", tool, status)
// 	}
// }

// // Parse output and extract tools and statuses
// func parseOutput(out string) map[string]string {
// 	lines := strings.Split(out, "\n")
// 	tools := make(map[string]string)

// 	// Start reading from the line with tool information
// 	startReading := false
// 	for _, line := range lines {
// 		if strings.HasPrefix(line, "1. ") {
// 			startReading = true
// 			continue
// 		}
// 		if startReading {
// 			// Split the line into tool and status
// 			parts := strings.SplitN(line, " ", 2)
// 			if len(parts) == 2 {
// 				tools[parts[1]] = parts[0]
// 			}
// 		}
// 	}

// 	return tools
// }

// ======================================================================================

/*
Version 2 of the above
- This is from GPT3 , and its much better
*/

// Output struct
type Output struct {
	Tools map[string]string
}

func V22() {
	// Execute PDTM command
	out, err := exec.Command("pdtm").Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Convert output to string
	output := string(out)

	// Parse output and extract tools and statuses
	tools := parseOutput(output)

	// Print tools and statuses with only the first two words
	for tool, status := range tools {
		// Split the tool into words
		words := strings.Split(tool, " ")
		if len(words) >= 2 {
			toolName := words[0]   // First word
			// toolStatus := words[1] // Second word
			fmt.Printf("Tool: %s, Status: %s\n", toolName, status)
		}
	}
}

// Parse output and extract tools and statuses
func parseOutput(out string) map[string]string {
	lines := strings.Split(out, "\n")
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
