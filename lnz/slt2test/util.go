/*
Using the curl command to print out contents of the  snips
https://snips.sh/f/yVqWKywdT2
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

/*
This will fetch the contets of https://snips.sh/f/yVqWKywdT2 and print
- This is exact copy of sl2/v2/pusi.go
*/
func U1() {
	// Defining the variables for command an porams
	progName := "curl"
	progArgs := "https://snips.sh/f/yVqWKywdT2"

	// Command to execute
	cmd := exec.Command(progName, progArgs)

	// Get output from command
	var out bytes.Buffer
	cmd.Stdout = &out

	// Execute command
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Print output
	fmt.Println(out.String())
}
