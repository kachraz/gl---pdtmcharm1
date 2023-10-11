/*
V1 of the work which you are doing , doing it this way for testing stuff which you have alreaady learn
*/

package work

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// -------------------------------------------------------------
// Application will start here

// This is the main caller of the function
func V1m() {
	fmt.Println("V1 of the work which you are doing , doing it this way for testing stuff which you have alreaady learn")
}

// ---------------------------------

// Declare the Model struct as per the documentation

type model struct {
	choices  []string         // items on the to-do list
	cursor   int              // which to-do list item our cursor is pointing at
	selected map[int]struct{} // which to-do items are selected
}

// This is the initial model function also  taken verbaitm from the document

func initialModel() model {
	return model{
		// Our to-do list is a grocery list
		choices: []string{"Fluffy Booty", "Hairy Smelly Pussy", "Sweaty Armpits"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

// Implementing the model interface - Also taken from the manual
// Note this it starndard method for this type

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

// This is the update model - Also taken from the manual

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+d", "q": // You cahnged 
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}
