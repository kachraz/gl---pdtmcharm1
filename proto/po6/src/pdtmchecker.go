/*
シ MXIUM SYSTEMS シ
pdtmchecker.go
Checks in pdtm in installed
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

var (
	programName = "pdtm"
	programArgs = " "
	s           = `
[HINT] go install -v github.com/projectdiscovery/pdtm/cmd/pdtm@latest
`
)

// Main Function
func PdtsChecker() {

	cr := C.New(C.FgRed).SprintFunc()
	chg := C.New(C.FgHiGreen).SprintFunc()

	cmd := exec.Command(programName, programArgs)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	logger := log.New(os.Stderr, "", 0)
	cmd.Stderr = logger.Writer()

	err := cmd.Run()

	if err != nil {
		fmt.Println(cr("-------------------------------------"))
		logger.Printf(cr("[FAIL] %s"), err)
		fmt.Println(cr("[FAIL] PDTM - Project Discovery Open Source Tool Manager - not installed!\n" + s))
		fmt.Println(cr("-------------------------------------"))
	} else {

		fmt.Println(chg("\n[OK] Ran PDTS Successfully!...Choose Install below..."))
		TableFuncMain()
	}

}
