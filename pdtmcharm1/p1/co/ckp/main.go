/*
This will check if a certain program is installed or not
*/

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	C "github.com/fatih/color"
)

// Main variable that represents the program
var (
	programName = "pdtm"
	programArgs = "--version"
	s = `
	go install -v github.com/projectdiscovery/pdtm/cmd/pdtm@latest
	`
)

func main() {

	// Color definitions
	cr := C.New(C.FgRed).SprintFunc()
	cg := C.New(C.FgGreen).SprintFunc()

	// Create a new logger with the custom color function
	logger := log.New(os.Stderr, "", 0)
	logger.SetFlags(0) // Remove default timestamp and file name flags
	logger.SetOutput(C.Output)

	cmd := exec.Command(programName, programArgs)
	err := cmd.Run()
	if err != nil {
		logger.Fatalf(cr("[FAIL] %v \n %v"), cr(err), cr(s) )
		fmt.Printf("Install - go get ....\n")
	}
	logger.Println(cg("[OK] Program is installed"))
}
