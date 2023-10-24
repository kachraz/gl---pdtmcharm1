/*
シ MXIUM SYSTEMS シ
main.go
Main Entry Point of the application
*/

package main

import (
	"flag"
	"fmt"
	B "po/src"
)

func main() {
	FuncHelp()
	B.Banmain()
	B.PdtsMain()
}

// Function help Text

func FuncHelp() {
	help := flag.Bool("help", false, "Display help message")
	flag.Parse()

	if *help {
		displayHelp()
		return
	}
}

func displayHelp() {
	fmt.Println("This is the help message.")
}
