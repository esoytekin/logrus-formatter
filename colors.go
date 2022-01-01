package formatter

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

// ANSI color codes.
const (
	AnsiReset     = 0
	AnsiRed       = 31
	AnsiHiRed     = 91
	AnsiGreen     = 32
	AnsiHiGreen   = 92
	AnsiYellow    = 33
	AnsiHiYellow  = 93
	AnsiBlue      = 34
	AnsiHiBlue    = 94
	AnsiMagenta   = 35
	AnsiHiMagenta = 95
	AnsiCyan      = 36
	AnsiHiCyan    = 96
	AnsiWhite     = 37
	AnsiHiWhite   = 97
)

// Color colorizes the input string and returns it with ANSI color codes.
func (formatter *CustomFormatter) Color(entry *logrus.Entry, s string) string {

	if !formatter.ForceColors && formatter.DisableColors {
		return s
	}

	// Determine color. Default is info.
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = formatter.ColorDebug
	case logrus.WarnLevel:
		levelColor = formatter.ColorWarn
	case logrus.ErrorLevel:
		levelColor = formatter.ColorError
	case logrus.PanicLevel:
		levelColor = formatter.ColorPanic
	case logrus.FatalLevel:
		levelColor = formatter.ColorFatal
	default:
		levelColor = formatter.ColorInfo
	}
	if levelColor == AnsiReset {
		return s
	}

	// Colorize.
	return "\033[" + strconv.Itoa(levelColor) + "m" + s + "\033[0m"
}
