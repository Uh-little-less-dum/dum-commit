package cli_styles

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// WARN: This is duplicated in the cli package. Consider creating a single utility go package to collect this shared content.
var UlldBlue = "#0ba5e9"

var DocStyle = lipgloss.NewStyle().Margin(0, 2)

// HelpStyle styling for help context menu
var HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

// ErrStyle provides styling for error messages
var ErrStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#bd534b")).Render

// AlertStyle provides styling for alert messages
var AlertStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("62")).Render

var MutedTextColor = lipgloss.Color("241")

var UlldBlueLipgloss = lipgloss.Color(UlldBlue)

var TitleStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFDF5")).
	Background(lipgloss.Color(UlldBlue)).
	Padding(0, 1)

var UlldBlueForeground = lipgloss.NewStyle().Foreground(UlldBlueLipgloss)

func GetHuhTheme() *huh.Theme {
	t := huh.ThemeCharm()
	t.Group.Padding(5).Margin(5)
	t.Blurred.Title = UlldBlueForeground
	t.Focused.Title = UlldBlueForeground
	t.Focused.FocusedButton = t.Blurred.FocusedButton.Background(UlldBlueLipgloss)
	t.Blurred.FocusedButton = t.Blurred.FocusedButton.Background(UlldBlueLipgloss)
	t.Blurred.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(MutedTextColor)
	t.Focused.TextInput.Prompt = t.Focused.TextInput.Prompt.Foreground(lipgloss.Color("#ffffff"))
	t.Focused.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(MutedTextColor)
	t.Blurred.TextInput.Cursor = t.Focused.TextInput.Cursor.Foreground(MutedTextColor)
	t.Focused.SelectedOption = t.Focused.SelectedOption.Foreground(UlldBlueLipgloss)
	t.Blurred.SelectedOption = t.Focused.SelectedOption.Foreground(MutedTextColor)
	// t.Focused.Base = lipgloss.NewStyle().PaddingLeft(1).BorderStyle(lipgloss.ThickBorder()).BorderLeft(true).BorderForeground(UlldBlueLipgloss)

	return t
}
