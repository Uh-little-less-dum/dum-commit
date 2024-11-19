package init_utils

import (
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

var UlldBlue = "#0ba5e9"

func setLoggerStyles() {
	styles := log.DefaultStyles()
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString("ERROR").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("204")).
		Foreground(lipgloss.Color("0"))

	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString("DEBUG").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color("204"))
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString("Info").
		Padding(0, 1, 0, 1).
		Foreground(lipgloss.Color(UlldBlue))

	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString("WARN").
		Padding(0, 2, 0, 1).
		Foreground(lipgloss.Color("#ffff00"))

	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString("Oh Shit...").
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color("204")).
		Foreground(lipgloss.Color("0"))
	styles.Keys["err"] = lipgloss.NewStyle().Foreground(lipgloss.Color("204"))
	styles.Values["err"] = lipgloss.NewStyle().Bold(true)
	log.SetStyles(styles)
}

func initLogger() {
	log.SetTimeFormat(time.Kitchen)
	verbose := viper.GetBool("verbose")
	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	setLoggerStyles()

}
