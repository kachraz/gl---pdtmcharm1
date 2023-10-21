/*
pdts.go
This file is for
1. Checking if pdts is installed or not
and then it will install it
*/

package src

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	C "github.com/fatih/color"
)

func PdtsMain() {
	PdtsChecker()
}

// Main variable that represents the program
var (
	programName = "pdtm"
	programArgs = "--version"
	s           = `
	go install -v github.com/projectdiscovery/pdtm/cmd/pdtm@latest
	`
)

func PdtsChecker() {

	// Color definitions
	cr := C.New(C.FgRed).SprintFunc()
	cg := C.New(C.FgGreen).SprintFunc()
	cb := C.New(C.FgBlue).SprintFunc()

	// Mention checking for pdts
	fmt.Println(cb(`Checking if Pdts is installed...`))

	// Create a new logger with the custom color function
	logger := log.New(os.Stderr, "", 0)
	logger.SetFlags(0) // Remove default timestamp and file name flags
	logger.SetOutput(C.Output)

	cmd := exec.Command(programName, programArgs)
	err := cmd.Run()
	if err != nil {
		logger.Fatalf(cr("[FAIL] %v \n %v"), cr(err), cr(s))
		fmt.Printf("Install - go get ....\n")
	}
	logger.Println(cg("[OK] Pdts is Installed"))
}
