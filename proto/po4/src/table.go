/*
table.go
This file has the actual table code
*/

package src

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func TableFuncMain() {
	TableMain()
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("200")) // This controls the overall border color

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			// This is your edit
			selectedItem := m.table.SelectedRow()[0] // This controls which item from the table is being selected starts with 0
			cmd := exec.Command("echo", selectedItem)
			// cmd := exec.Command("pwd", selectedItem)
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("Fucker bastard !", err)
				return m, tea.Quit
			}
			fmt.Println(string(output))
			return m, tea.Quit
			// return m, tea.Batch(
			// 	tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			// )
			// till here ------------
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// ---------------------------- Your Edits Start --------------------------
// ---------------------------- Your Edits Start --------------------------
// ---------------------------- Your Edits Start --------------------------
/*
This is ou addition of the help functions taken from gla ,
adding the navigation help at the end of the table
*/
var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("200")).Render

func (m model) helpView() string {
	return helpStyle("\n  ↑/↓: Navigate • q: Quit • Enter: to Select Install\n")
}

// ---------------------------- Your Edits ends --------------------------
// ---------------------------- Your Edits ends --------------------------
// ---------------------------- Your Edits ends --------------------------

// This controls the views of the application
// helpView() added here
func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n" + m.helpView() + "\n"
}

// This is the main program and entry poin , this will go up in V1main()
func TableMain() {
	columns := []table.Column{
		{Title: "Rank", Width: 4},
		{Title: "City", Width: 10},
		{Title: "Country", Width: 10},
		{Title: "Population", Width: 20},
	}

	rows := []table.Row{
		{"1", "aix", "AIx is a cli tool to interact with Large Language Models (LLM) APIs.", "github.com/projectdiscovery/aix"},
		{"2", "alterx", "Fast and customizable subdomain wordlist generator using DSL.", "github.com/projectdiscovery/alterx"},
		{"3", "asnmap", "Go CLI and Library for quickly mapping organization network ranges using", "github.com/projectdiscovery/asnmap"},
		{"4", "cdncheck", "Tool for identifying the technology associated with dns / ip network addresses.", "github.com/projectdiscovery/cdncheck"},
		{"5", "chaos-client", "Go client to communicate with Chaos dataset API.", "github.com/projectdiscovery/chaos-client"},
		{"6", "cloudlist", "multi-cloud tool for getting Assets from Cloud Providers", "github.com/projectdiscovery/cloudlist"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10), // This controls number of rows
	)

	/*
		This section is the style it follows the ANSI 256 Color Codes
		https://www.ditig.com/256-colors-cheat-sheet
	*/
	s := table.DefaultStyles()
	s.Header = s.Header. // Affects header only
				BorderStyle(lipgloss.NormalBorder()).
				Foreground(lipgloss.Color("229")).
				BorderForeground(lipgloss.Color("82")).
				BorderBottom(true).
				Bold(false)
	s.Selected = s.Selected. // Affects Selections only
					Foreground(lipgloss.Color("229")).
					Background(lipgloss.Color("22")).
					Bold(false).
					Italic(true) // Added Italics

	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
