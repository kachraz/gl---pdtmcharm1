/*
シ MXIUM SYSTEMS シ
main.go
Main Entry Point of the application
*/

package main

import (
	"fmt"
	"os"
	B "po/src"
)

func main() {

	// When help flags executed

	if len(os.Args) > 1 && os.Args[1] == "--help" {
		printHelp()
		return
	}

	// Main Funtion entry Point
	B.Banmain()
	B.PdtsMain()
}

// Help Function
func printHelp() {
	fmt.Println("This program shows purchase order details")
	fmt.Println("Usage:")
	fmt.Println("  --help   Print this help")
}
