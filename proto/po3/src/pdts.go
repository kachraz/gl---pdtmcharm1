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
	programName = "pdto"
	programArgs = " "
	s           = `
go install -v github.com/projectdiscovery/pdtm/cmd/pdtm@latest
	`
	Ban = `
	
                ____          
     ____  ____/ / /_____ ___ 
    / __ \/ __  / __/ __ __  \
   / /_/ / /_/ / /_/ / / / / /
  / .___/\__,_/\__/_/ /_/ /_/ 
 /_/                         

                projectdiscovery.io
	`
)

// From Claude and Codeium
func PdtsChecker() {

	// Print header of program
	// fmt.Println(Ban)

	//Color Definitions
	cr := C.New(C.FgRed).SprintFunc()
	chg := C.New(C.FgHiGreen).SprintFunc()

	// pdts command
	cmd := exec.Command(programName, programArgs)

	// Set stdout to print output
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr // This also prints out the stderr

	// Settingup the logger here to print logger message also
	logger := log.New(os.Stderr, "", 0)
	cmd.Stderr = logger.Writer()

	// Run command
	err := cmd.Run()

	if err != nil {
		// pdts not installed
		logger.Printf(cr("[FAIL] %s"), err)
		fmt.Println(cr("[FAIL] PDTM - Project Discovery Open Source Tool Manager - not installed!" + s))
	} else {

		// Print success
		fmt.Println(chg("\n[OK] Ran PDTS Successfully!...Choose Install below..."))
		TableFuncMain()
	}

}
