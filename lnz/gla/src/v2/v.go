/*
Making the same glamor program as in v1 , but with a viewport which is describe here
https://github1s.com/charmbracelet/bubbletea/blob/HEAD/examples/glamour/main.go
*/

package src

import (
	"fmt"
	"os"
	v1 "gl/src/v1"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

func V2main() {
	V2ProgMain()
}

// This section it taking all the code from the example and changing function names

const SmellyPanty = `
# Today’s Stinky Fetishes

## Appetizers

> Smelly unwiped ass and pussy specials

| Name        | Price | Notes                           |
| ---         | ---   | ---                             |
| Sweaty Pussy   | $2    | Just an appetizer               |
| Ass Juice | $4    | Made with San Marzano tomatoes  |


## Desserts

| Name         | Price | Notes                 |
| ---          | ---   | ---                   |
| Squirt     | $4    | Looks good on rabbits |
| Milk Enema | $5    | A classic             |
| Pussy Cream   | $3    | Pretty creamy!        |

Made my drity strict bbw scatqueens 

* [x] Mistress ShitOnYou
* [x] Mistress ForciblyFartInNose
* [ ] Mistress SweatNShot

Bon appétit!
`

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

type v2 struct {
	viewport viewport.Model
}

func newGlam() (*v2, error) {
	const width = 78

	vp := viewport.New(width, 20)
	vp.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		PaddingRight(2)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return nil, err
	}

	str, err := renderer.Render(SmellyPanty)
	if err != nil {
		return nil, err
	}

	vp.SetContent(str)

	return &v2{
		viewport: vp,
	}, nil
}

func (v v2) Init() tea.Cmd {
	v1.AssPrint()
	return nil
}

func (v v2) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			fmt.Println("\n\nFuck off bastard\n\n")
			return v, tea.Quit
		default:
			var cmd tea.Cmd
			v.viewport, cmd = v.viewport.Update(msg)
			return v, cmd
		}
	default:
		return v, nil
	}
}

func (v v2) View() string {
	return v.viewport.View() + v.helpView()
}

func (v v2) helpView() string {
	return helpStyle("\n  ↑/↓: Navigate • q: Quit\n")
}

func V2ProgMain() {
	model, err := newGlam()
	if err != nil {
		fmt.Println("Could not initialize Bubble Tea model:", err)
		os.Exit(1)
	}

	if _, err := tea.NewProgram(model).Run(); err != nil {
		fmt.Println("Bummer, there's been an error:", err)
		os.Exit(1)
	}
}
