// Attempting to make a lipglo function
/*

_________________________$$$$$$$________________
________________________$$$$$$$$$$______________
________________________$$$$$$$$$$$_____________
_________________________$$$$$$$$$$$$$$_________
__________________________$$$$$$$$$$$___________
_____________________________$$$$$$$$$$$$$______
___________________________$$$$$$$$$$___________
_________________________$$$$$$$$$$$$$$$________
________________$$$______$$$$$$$$$$$$$$_________
______________$$$$$$$$_____$$$$$$__$$$$$________
_____________$$$$$$$$$$_____$$$$____$$$$$_______
___________$$$$$$_$$$$$$$$__$$$$______$$$$______
__________$$$$$_____$$$$$$$$_$$$$_______$$$_____
________$$$$$_________$$$$$$$$$$$$_______$$$____
_______$$$_____________$$$$$$$$$$$________$$$___
_____$$$________________$$$$$$$$$$________$$$$$$
__$$$$$$__________________$$$$$$$_______________

main Functions to use 
1. Banr() - banner function
2. Headr() - header function
3. Texc() - text function
4. Para() - paragraph function
5. Kode() - code function
6. T1() - simple color string
7. T2() - simple color string
8. T3() - simple color string

*/


package lib

import (
	f "fmt"
	"time"

	lg "github.com/charmbracelet/lipgloss"
)

/*
The following fnction is super sexy
1. First of all you are making a simple function that has a receiver of type string
2. Then you are doing all the coloring with a lipgloss style
3. Note that since you put a receiver of type string, you are passing in a string
4. Its part of packge main , so you can now call it in any file amazingly
5. Then you just call it in any function ensuring that what is being passed in is a string
6. You can now call this function anywhere
sna
*/

// This is called a block creation
var (
	// Variable for banner function Banr()
	liq = lg.NewStyle().
		Foreground(lg.Color("#16FF00")).
		Background(lg.Color("#2D033B")).
		Width(100).Align(lg.Center).Bold(true).
		Blink(true).Border(lg.DoubleBorder(), true, true, true, true).
		PaddingTop(1).
		PaddingBottom(1).
		BorderForeground(lg.Color("#FC2947"))

		// Replace banner bbw ass spread licking here 
	banner = `
WRITE SHIT HERE 
`

	// Style for Headr()
	lqh = lg.NewStyle().
		Width(100).Align(lg.Center).Bold(true).
		Background(lg.Color("#0C1E7F")).Foreground(lg.Color("#FF008E")).
		Bold(true).Underline(true).
		PaddingTop(1).PaddingBottom(1)

		// Style for Texc()
	texco = lg.NewStyle().
		Width(100).Align(lg.Right).
		Background(lg.Color("#3A1078")).
		Foreground(lg.Color("#00DFA2"))

		// Style for Para()
	pa = lg.NewStyle().
		Width(100).
		Background(lg.Color("#2E4F4F")).
		Foreground(lg.Color("#ACFADF"))

		// Style for Kode()
	ko = lg.NewStyle().
		Foreground(lg.Color("#16FF00")).
		Italic(true)

		// Style for T1()
	t1 = lg.NewStyle().
		Foreground(lg.Color("#16FF00"))
		// Style for T2()
	t2 = lg.NewStyle().
		Foreground(lg.Color("#FF008E"))
		// Style for T3()
	t3 = lg.NewStyle().
		Foreground(lg.Color("#00DFA2"))
)

func Banr(s string) {
	t := time.Now().UTC().Format("Mon Jan 02 15:04:05 UTC")

	f.Println(liq.Render(banner, t, "\n\n", s))

}

// Thus funtion is for the nice heading colors
func Headr(s string) {
	f.Println("\n")
	f.Println(lqh.Render(s), "\n")
}

// This is for the sub header
func Texc(s string) {
	f.Println(texco.Render(s))
}

// This for the paragraph
func Para(s string) {
	f.Println(pa.Render(s), "\n")
}

// This is for the code
func Kode(s string) {
	f.Println(ko.Render(s))
}

// These are the simple color strings
func T1(s string) {
	f.Println(t1.Render(s))
}

func T2(s string) {
	f.Println(t2.Render(s))
}

func T3(s string) {
	f.Println(t3.Render(s))
}
