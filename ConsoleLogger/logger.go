package consolelogger

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

// ANSI escape codes for rainbow colors
var colors = []string{
	"\033[31m", // Red
	"\033[33m", // Yellow
	"\033[32m", // Green
	"\033[36m", // Cyan
	"\033[34m", // Blue
	"\033[35m", // Magenta
}

// Reset ANSI code
const reset = "\033[0m"

// rainbowColorize applies rainbow colors to each character of the input string
func rainbowColorize(input string) string {
	colorized := ""
	colorIndex := 0

	for _, char := range input {
		if char == '\n' {
			// Preserve newlines
			colorized += string(char)
		} else {
			// Apply color
			colorized += colors[colorIndex%len(colors)] + string(char) + reset
			colorIndex++
		}
	}

	return colorized
}

func PrintWelcomeMessage() {
	ascii := figure.NewFigure("SYNAPSE", "colossal", true).String()
	colorizedASCII := rainbowColorize(ascii)
	fmt.Println(colorizedASCII)
}

func InfoLog(message string) {
	fmt.Println(colors[4] + "INFO: " + message + reset)
}

func DebugLog(message string) {
	fmt.Println(colors[0] + "DEBUG: " + message + reset)
}
