/*
This go file is just for demonstrating the colors usage, which you can import as and when you need it
*/

package lib

import (
	f "fmt"

	C "github.com/fatih/color"
)

// The following function shouldnt be used anywhere ,its just for illustration of using colors
func FatihColorsShow() {

	// Defining all the fatih colors
	cr := C.New(C.FgRed).SprintFunc()
	chr := C.New(C.FgHiRed).SprintFunc()
	chrb := C.New(C.FgHiRed, C.Bold).SprintFunc()
	cg := C.New(C.FgGreen).SprintFunc()
	chg := C.New(C.FgHiGreen).SprintFunc()
	chgbu := C.New(C.FgHiGreen, C.Bold, C.Underline).SprintFunc()
	cy := C.New(C.FgYellow).SprintFunc()
	chy := C.New(C.FgHiYellow).SprintFunc()
	chyb := C.New(C.FgHiYellow, C.BgRed).SprintFunc()
	cb := C.New(C.FgBlue).SprintFunc()
	chb := C.New(C.FgHiBlue).SprintFunc()
	chbi := C.New(C.FgHiBlue, C.Italic).SprintFunc()
	cm := C.New(C.FgMagenta).SprintFunc()
	chm := C.New(C.FgHiMagenta).SprintFunc()
	chmc := C.New(C.FgHiMagenta, C.Concealed).SprintFunc()
	ccy := C.New(C.FgCyan).SprintFunc()
	chcy := C.New(C.FgHiCyan).SprintFunc()
	chcyr := C.New(C.FgHiCyan, C.ReverseVideo).SprintFunc()
	cw := C.New(C.FgWhite).SprintFunc()
	chw := C.New(C.FgHiWhite).SprintFunc()
	chwc := C.New(C.FgHiWhite, C.CrossedOut).SprintFunc()

	// Just print using all to remove errors shit
	// There is also bold and italic and you cna declare background also

	f.Printf(cr("This Red 			:%v\n"), cr("Red"))
	f.Printf(chr("This HIRed 		:%v\n"), chr("HiRed"))
	f.Printf(chrb("This HIRed 		:%v\n"), chrb("HiRedBold"))
	f.Printf(cg("This Green  		:%v\n"), cg("Green"))
	f.Printf(chg("This HiGreen  	:%v\n"), chg("HiGreen"))
	f.Printf(chgbu("This HiGreen Bold Underline  	:%v\n"), chgbu("HiGreen"))
	f.Printf(cy("This Yellow 		:%v\n"), cy("Yellow"))
	f.Printf(chy("This HiYellow 	:%v\n"), chy("HiYellow"))
	f.Printf(chyb("HiYellow + Background Red 	:%v\n"), chyb("HiYellow"))
	f.Printf(cb("This Blue 			:%v\n"), cb("Blue"))
	f.Printf(chb("This HiBlue 		:%v\n"), chb("HiBlue"))
	f.Printf(chbi("This HiBlue Italic 		:%v\n"), chbi("HiBlue"))
	f.Printf(cm("This Magenta 		:%v\n"), cm("Magenta"))
	f.Printf(chm("This HiMagenta  	:%v\n"), chm("HiMagenta"))
	f.Printf(chmc("This HiMagenta Concealed  	:%v\n"), chmc("HiMagenta Concealed"))
	f.Printf(cy("This Yellow  		:%v\n"), cy("Yellow"))
	f.Printf(ccy("This Cyan 		:%v\n"), ccy("Cyan"))
	f.Printf(chcy("This HiCyan  	:%v\n"), chcy("HiCyan"))
	f.Printf(chcyr("This HiCyan Reverse Video  	:%v\n"), chcyr("HiCyan"))
	f.Printf(cw("This White  		:%v\n"), cw("White"))
	f.Printf(chw("This HiWhite  	:%v\n"), chw("HiWhite"))
	f.Printf(chwc("This HiWhite CrossedOut 	:%v\n"), chwc("HiWhite"))

}

// Second style from above putting single words
func Coline() {

	//Color Definitions

	cr := C.New(C.FgRed).SprintFunc()
	chr := C.New(C.FgHiRed).SprintFunc()
	chrb := C.New(C.FgHiRed, C.Bold).SprintFunc()
	cg := C.New(C.FgGreen).SprintFunc()
	chg := C.New(C.FgHiGreen).SprintFunc()
	chgbu := C.New(C.FgHiGreen, C.Bold, C.Underline).SprintFunc()
	cy := C.New(C.FgYellow).SprintFunc()
	chy := C.New(C.FgHiYellow).SprintFunc()
	chyb := C.New(C.FgHiYellow, C.BgRed).SprintFunc()
	cb := C.New(C.FgBlue).SprintFunc()
	chb := C.New(C.FgHiBlue).SprintFunc()
	chbi := C.New(C.FgHiBlue, C.Italic).SprintFunc()
	cm := C.New(C.FgMagenta).SprintFunc()
	chm := C.New(C.FgHiMagenta).SprintFunc()
	chmc := C.New(C.FgHiMagenta, C.Concealed).SprintFunc()
	ccy := C.New(C.FgCyan).SprintFunc()
	chcy := C.New(C.FgHiCyan).SprintFunc()
	chcyr := C.New(C.FgHiCyan, C.ReverseVideo).SprintFunc()
	cw := C.New(C.FgWhite).SprintFunc()
	chw := C.New(C.FgHiWhite).SprintFunc()
	chwc := C.New(C.FgHiWhite, C.CrossedOut).SprintFunc()

	f.Println(cr("Red"), chr("HiRed"), chrb("HiRedBold"), cg("Green"), chg("HiGreen"), chgbu("HiGreen"), cy("Yellow"), chy("HiYellow"),"\n" ,chyb("HiYellow"), cb("Blue"), chb("HiBlue"), chbi("HiBlue"), cm("Magenta"), chm("HiMagenta"), chmc("HiMagenta Concealed"),"\n", cy("Yellow"), ccy("Cyan"), chcy("HiCyan"), chcyr("HiCyan"), cw("White"), chw("HiWhite"), chwc("HiWhite"))

}

// This was used above , keeping her for reference
// cr("Red")),chr("HiRed")),chrb("HiRedBold")),cg("Green")),chg("HiGreen")),chgbu("HiGreen")),cy("Yellow")),chy("HiYellow")),chyb("HiYellow")),cb("Blue")),chb("HiBlue")),chbi("HiBlue")),cm("Magenta")),chm("HiMagenta")),chmc("HiMagenta Concealed")),cy("Yellow")),ccy(("Cyan")),chcy("HiCyan")),chcyr("HiCyan")),cw("White")),chw("HiWhite")),chwc("HiWhite"))