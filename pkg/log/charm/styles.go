package log

import (
	"github.com/Genekkion/gogogadgets/pkg/log"
	"github.com/charmbracelet/lipgloss"
	cl "github.com/charmbracelet/log"
)

var defaultStyle = func() *cl.Styles {
	styles := cl.DefaultStyles()

	style := lipgloss.NewStyle().
		Bold(true).
		Padding(0, 1, 0, 1)

	styles.Levels[cl.DebugLevel] = style.
		SetString(log.LevelDebug.Tag()).
		Background(lipgloss.Color("#414868")).
		Foreground(lipgloss.Color("#FFFFFF"))

	styles.Levels[cl.InfoLevel] = style.
		SetString(log.LevelInfo.Tag()).
		Background(lipgloss.Color("#485E30")).
		Foreground(lipgloss.Color("#FFFFFF"))

	styles.Levels[cl.WarnLevel] = style.
		SetString(log.LevelWarn.Tag()).
		Background(lipgloss.Color("#FF9E64")).
		Foreground(lipgloss.Color("#000000"))

	styles.Levels[cl.ErrorLevel] = style.
		SetString(log.LevelError.Tag()).
		Background(lipgloss.Color("#F7768E")).
		Foreground(lipgloss.Color("#000000"))

	styles.Levels[cl.FatalLevel] = style.
		SetString(log.LevelFatal.Tag()).
		Background(lipgloss.Color("#BB9AF7")).
		Foreground(lipgloss.Color("#000000"))

	return styles
}()
