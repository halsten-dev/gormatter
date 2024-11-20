package gormatter

import (
	"fmt"
	"strconv"
)

const (
	styleStart = "\033["
	styleEnd   = "m"
	styleReset = "\033[0m"
)

const (
	styleBold      = 1
	styleItalic    = 3
	styleUnderline = 4
)

const (
	foregroundMod     = 30
	backgroundMod     = 40
	foregroundHighMod = 90
	backgroundHighMod = 100
)

const None = -1

const (
	Black = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White

	HighBlack   = Black + 10
	HighRed     = Red + 10
	HighGreen   = Green + 10
	HighYellow  = Yellow + 10
	HighBlue    = Blue + 10
	HighMagenta = Magenta + 10
	HighCyan    = Cyan + 10
	HighWhite   = White + 10
)

func Normal(foregroundColor int, backgroundColor int, text string) string {
	style := styleStart
	style += colors(foregroundColor, backgroundColor)
	style += styleEnd

	return fmt.Sprintf("%s%s%s", style, text, styleReset)
}

func Bold(foregroundColor int, backgroundColor int, text string) string {
	style := styleStart
	style += attr(styleBold, true)
	style += colors(foregroundColor, backgroundColor)
	style += styleEnd

	return fmt.Sprintf("%s%s%s", style, text, styleReset)
}

func Italic(foregroundColor int, backgroundColor int, text string) string {
	style := styleStart
	style += attr(styleItalic, true)
	style += colors(foregroundColor, backgroundColor)
	style += styleEnd

	return fmt.Sprintf("%s%s%s", style, text, styleReset)
}

func Underline(foregroundColor int, backgroundColor int, text string) string {
	style := styleStart
	style += attr(styleUnderline, true)
	style += colors(foregroundColor, backgroundColor)
	style += styleEnd

	return fmt.Sprintf("%s%s%s", style, text, styleReset)
}

func BoldItalic(foregroundColor int, backgroundColor int, text string) string {
	style := styleStart
	style += attr(styleBold, true)
	style += attr(styleItalic, true)
	style += colors(foregroundColor, backgroundColor)
	style += styleEnd

	return fmt.Sprintf("%s%s%s", style, text, styleReset)
}

func BoldUnderline(foregroundColor int, backgroundColor int, text string) string {
	style := styleStart
	style += attr(styleBold, true)
	style += attr(styleUnderline, true)
	style += colors(foregroundColor, backgroundColor)
	style += styleEnd

	return fmt.Sprintf("%s%s%s", style, text, styleReset)
}

func BoldItalicUnderline(foregroundColor int, backgroundColor int, text string) string {
	style := styleStart
	style += attr(styleBold, true)
	style += attr(styleItalic, true)
	style += attr(styleUnderline, true)
	style += colors(foregroundColor, backgroundColor)
	style += styleEnd

	return fmt.Sprintf("%s%s%s", style, text, styleReset)
}

func ItalicUnderline(foregroundColor int, backgroundColor int, text string) string {
	style := styleStart
	style += attr(styleItalic, true)
	style += attr(styleUnderline, true)
	style += colors(foregroundColor, backgroundColor)
	style += styleEnd

	return fmt.Sprintf("%s%s%s", style, text, styleReset)
}

func colors(foregroundColor, backgroundColor int) string {
	var style string

	if foregroundColor != None {
		style += attr(fgColor(foregroundColor), false)
	}

	if backgroundColor != None {
		if len(style) > 0 {
			style += ";"
		}

		style += attr(bgColor(backgroundColor), false)
	}

	return style
}

func attr(a int, addSeparator bool) string {
	attr := strconv.Itoa(a)

	if addSeparator {
		attr += ";"
	}

	return attr
}

func fgColor(color int) int {
	highColor := color >= 10

	if highColor {
		return foregroundHighMod + (color - 10)
	}

	return foregroundMod + color
}

func bgColor(color int) int {
	highColor := color >= 10

	if highColor {
		return backgroundHighMod + (color - 10)
	}

	return backgroundMod + color
}
