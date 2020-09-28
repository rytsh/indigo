package common

import (
	"os"

	"github.com/fatih/color"
)

// Color is hold color function
var Color = map[string]*color.Color{
	"Bold": color.New(color.Bold),

	"Magenta": color.New(color.Reset, color.FgMagenta),
	"Red":     color.New(color.FgHiRed),
	"Yellow":  color.New(color.FgHiYellow),
	"Blue":    color.New(color.FgHiBlue),
	"Green":   color.New(color.FgHiGreen),

	"Error": color.New(color.FgHiWhite, color.BgRed),
	"Warn":  color.New(color.FgHiWhite, color.BgYellow),

	"Reset": color.New(color.Reset),
}

// PrintIntro is printing Intro
func PrintIntro() {
	Color["Bold"].Println(Intro)
}

// DisableColor for without color output
func DisableColor() {
	color.NoColor = true
}

// ErrorPrintExit is block red line print and exit
func ErrorPrintExit(msg string, exit int) {
	Color["Error"].Print("\n[ERR]: ", msg)
	Color["Reset"].Println()
	os.Exit(2)
}
