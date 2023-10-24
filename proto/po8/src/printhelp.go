/*
シ MXIUM SYSTEMS シ
printhelp.go
Prints the help test
*/

package src

import "fmt"

func FuncHelp() {
	Banmain()
	var helpText = `
...................,'...................
...............'...;:lll'...............
..............'..'.,,:;lxx'.............
.............; .'..,,:;,cdO'............    
............:, ... '';,,:lkc............    
...........,'..       .';ldd............	Function calls PDTM andit has to be installed in order for it to work
...........:.'   ....':,.';:;...........	made by - https://github.com/m0ham3dx
............;,  . .':..;.';,:...........
...........,'.......:,;:',c.,..........'	Helper function for PDTM that displays a menu with desriptions of tools
...........'.,.'.  .,'....:..:......''''
............... ,.  ........;;;c'''.....	Function calls PDTM andit has to be installed in order for it to work
............    ..  ..... .;':.l:',,''''
.........';      .. .. ..''','';clc;,,'.
.......,'..,.       .......'';ccllddlc;'
.....,,.   '...   .......;l::c::;;;,;:lo
.,;:,. .    ..............;llc;,,,,;;;cl
.,.....:..       ....';;,;,'.,,'',:::;'.
;.,....'.'.  ....'..,,;:;,',,...,,;''.;l
.... .. ... . ........'.',,'. ..'....';c
  .    ..... ...........'.'.. .....  ...
	`
	fmt.Println(helpText)
}
