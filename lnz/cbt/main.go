/*
Main entry point of the application
*/

package main

import (
	A "cbt/lib"
	V1 "cbt/ver/v1"
	V2 "cbt/ver/v2"
	"fmt"
)

func main() {
	A.Asprint()
	fmt.Println("Hello World")
	V1.MainV1()
	V2.MainV2()
}
