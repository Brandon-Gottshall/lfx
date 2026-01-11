package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))
	errorStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("9"))
	warningStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("214"))
)

func PrintTitle(text string) {
	fmt.Println(titleStyle.Render(text))
}

func PrintError(message string, err error) {
	fmt.Fprintln(os.Stderr, errorStyle.Render("error:"), message)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func PrintWarning(message string, err error) {
	fmt.Fprintln(os.Stderr, warningStyle.Render("warning:"), message)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
