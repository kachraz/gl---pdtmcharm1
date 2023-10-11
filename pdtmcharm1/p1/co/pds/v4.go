/*
v4 - This is from codeium
*/

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func V4() {
	// Execute the pdtm command
	cmd := exec.Command("pdtm")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Extract the tool and its status from the output
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, " ") {
			continue
		}
		parts := strings.Split(line, " ")
		if len(parts) < 2 {
			continue
		}
		tool := parts[0]
		status := parts[1]
		fmt.Println(tool, status) // Print only the tool name and status
	}
}
