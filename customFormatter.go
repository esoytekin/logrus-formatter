package formatter

type CustomFormatter struct {
	// Set to true to bypass checking for a TTY before outputting colors.
	ForceColors bool

	// Force disabling colors and bypass checking for a TTY.
	DisableColors bool

	// Different colors for different log levels.
	ColorDebug int
	ColorInfo  int
	ColorWarn  int
	ColorError int
	ColorFatal int
	ColorPanic int
}

// NewFormatter creates a new CustomFormatter, sets the Template string, and returns its pointer.
// This function is usually called just once during a running program's lifetime.
//
// :param template: Pre-processed formatting template (e.g. "%[message]s\n").
//
// :param custom: User-defined formatters evaluated before built-in formatters. Keys are attributes to look for in the
// 	formatting string (e.g. "%[myFormatter]s") and values are formatting functions.
func NewFormatter() *CustomFormatter {
	formatter := CustomFormatter{
		ColorDebug: AnsiCyan,
		ColorInfo:  AnsiGreen,
		ColorWarn:  AnsiYellow,
		ColorError: AnsiRed,
		ColorFatal: AnsiMagenta,
		ColorPanic: AnsiMagenta,
	}

	return &formatter
}
