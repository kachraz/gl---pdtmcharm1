/*
Version 1 of testing being done here
*/

package src

import (
	"fmt"

	G "github.com/charmbracelet/glamour"
)

// Main function starts here

func Vmain() {

	// Defining the markdown file
	mk := `# Testing out Glamor File
- This shit better 
1. Look good
	1. Or else doggie 

~~~go 
package rapist 
import "fmt"
func rape() {
	Drink her piss
}
~~~

# NoNo 

> Also writing this code lets see

## How to Lick 

> Fuck all night
or else doggie

`

	out, err := G.Render(mk, "dark")
	if err != nil {
		// Handle the error here
		fmt.Println("Error:", err)
	} else {
		fmt.Print(out)
	}
}
