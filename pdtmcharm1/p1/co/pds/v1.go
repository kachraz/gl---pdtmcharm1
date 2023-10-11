/*
Actual file that is going to get the methods
- These methods suggested from claude
*/

package main

import (
	"fmt"
	"log"
	"os/exec"
)

/*
Method 1 of capturing output - Testing for the ls command=
*/
func V11() {

	out, err := exec.Command("ls", "-l").CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	formatted := fmt.Sprintf("OUTPUT:\n%s\n", string(out))
	fmt.Print(formatted)
}

/*
Idiomatic way of capturing the command , this is usin variadics
*/

func runCommand(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	return cmd.CombinedOutput()
}

func V12() {
	out, err := runCommand("pdtm")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", out)
}

/*
Method 3 - Again asking Claude the best way of doing it for the ls command ,
and capturing it with maps , Maps is better since iteration is not needed
and you can directly access the output
*/
// Output struct
// type Output struct {
// 	Header  string
// 	Divider string
// 	Tools   map[string]string
// }

// func Met3old() {

// 	// Execute command
// 	out, err := exec.Command("pdtm").Output()
// 	if err != nil {
// 		// Return error instead of printing
// 		return fmt.Errorf("error executing pdtm: %w", err)
// 	}

// 	// Convert output to string
// 	output := string(out)

// 	// Parse output
// 	o, err := parseOutput(output)
// 	if err != nil {
// 		return err
// 	}

// 	// Print tools
// 	if err := printTools(o); err != nil {
// 		return err
// 	}

// }

// // Parse output into structured map
// func parseOutput(out string) (Output, error) {

// 	lines := strings.Split(out, "\n")

// 	if len(lines) < 8 {
// 		return Output{}, fmt.Errorf("invalid output")
// 	}

// 	o := Output{
// 		Header:  strings.Join(lines[0:6], "\n"),
// 		Divider: lines[7],
// 		Tools:   make(map[string]string),
// 	}

// 	// Populate tools map
// 	for _, l := range lines[8:] {
// 		parts := strings.SplitN(l, " ", 2)
// 		if len(parts) < 2 {
// 			return Output{}, fmt.Errorf("invalid tool line: %s", l)
// 		}
// 		o.Tools[parts[0]] = parts[1]
// 	}

// 	return o, nil
// }

// // Print tools from parsed map
// func printTools(o Output) error {

// 	for tool, status := range o.Tools {
// 		_, err := fmt.Printf("%s %s\n", tool, status)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
