/*
pdts.go
This file is for
1. Checking if pdts is installed or not
and then it will install it
*/

package src

import (
	"fmt"
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
	programArgs = " "
	s           = `
	go install -v github.com/projectdiscovery/pdtm/cmd/pdtm@latest
	`
)

// Check if pdts is installed and then run it
// func PdtsChecker() {

// 	// Color definitions
// 	cr := C.New(C.FgRed).SprintFunc()
// 	cg := C.New(C.FgGreen).SprintFunc()

// 	// Create a new logger with the custom color function
// 	logger := log.New(os.Stderr, "", 0)
// 	logger.SetFlags(0) // Remove default timestamp and file name flags
// 	logger.SetOutput(C.Output)

// 	cmd := exec.Command(programName, programArgs)
// 	err := cmd.Run()
// 	if err != nil {
// 		logger.Fatalf(cr("[FAIL] %v \n %v"), cr(err), cr(s))
// 		fmt.Printf("Install - go get ....\n")
// 	}
// 	logger.Println(cg("[OK] Program is installed"))
// }

// From claude
func PdtsChecker() {

	//Color Definitions
	cr := C.New(C.FgRed).SprintFunc()
	chg := C.New(C.FgHiGreen).SprintFunc()

	// pdts command
	cmd := exec.Command(programName, programArgs)

	// Set stdout to print output
	cmd.Stdout = os.Stdout

	// Run command
	err := cmd.Run()

	if err != nil {
		// pdts not installed
		fmt.Println(cr("Pdts not installed!"))
	} else {

		// Print success
		fmt.Println(chg("\nRan pdts successfully!\n ---"))
	}

}
