package messages

import "github.com/logrusorgru/aurora/v4"

func Reset(arg any) aurora.Value {
	return c.Reset(arg)
}

func Clear(arg any) aurora.Value {
	return c.Clear(arg)
}

func Bold(arg any) aurora.Value {
	return c.Bold(arg)
}

func Faint(arg any) aurora.Value {
	return c.Faint(arg)
}

func DoublyUnderline(arg any) aurora.Value {
	return c.DoublyUnderline(arg)
}

func Fraktur(arg any) aurora.Value {
	return c.Fraktur(arg)
}

func Italic(arg any) aurora.Value {
	return c.Italic(arg)
}

func Underline(arg any) aurora.Value {
	return c.Underline(arg)
}

func SlowBlink(arg any) aurora.Value {
	return c.SlowBlink(arg)
}

func RapidBlink(arg any) aurora.Value {
	return c.RapidBlink(arg)
}

func Blink(arg any) aurora.Value {
	return c.Blink(arg)
}

func Reverse(arg any) aurora.Value {
	return c.Reverse(arg)
}

func Inverse(arg any) aurora.Value {
	return c.Inverse(arg)
}

func Conceal(arg any) aurora.Value {
	return c.Conceal(arg)
}

func Hidden(arg any) aurora.Value {
	return c.Hidden(arg)
}

func CrossedOut(arg any) aurora.Value {
	return c.CrossedOut(arg)
}

func StrikeThrough(arg any) aurora.Value {
	return c.StrikeThrough(arg)
}

func Framed(arg any) aurora.Value {
	return c.Framed(arg)
}

func Encircled(arg any) aurora.Value {
	return c.Encircled(arg)
}

func Overlined(arg any) aurora.Value {
	return c.Overlined(arg)
}

func Black(arg any) aurora.Value {
	return c.Black(arg)
}

func Red(arg any) aurora.Value {
	return c.Red(arg)
}

func Green(arg any) aurora.Value {
	return c.Green(arg)
}

func Yellow(arg any) aurora.Value {
	return c.Yellow(arg)
}

func Blue(arg any) aurora.Value {
	return c.Blue(arg)
}

func Magenta(arg any) aurora.Value {
	return c.Magenta(arg)
}

func Cyan(arg any) aurora.Value {
	return c.Cyan(arg)
}

func White(arg any) aurora.Value {
	return c.White(arg)
}

func BrightBlack(arg any) aurora.Value {
	return c.BrightBlack(arg)
}

func BrightRed(arg any) aurora.Value {
	return c.BrightRed(arg)
}

func BrightGreen(arg any) aurora.Value {
	return c.BrightGreen(arg)
}

func BrightYellow(arg any) aurora.Value {
	return c.BrightYellow(arg)
}

func BrightBlue(arg any) aurora.Value {
	return c.BrightBlue(arg)
}

func BrightMagenta(arg any) aurora.Value {
	return c.BrightMagenta(arg)
}

func BrightCyan(arg any) aurora.Value {
	return c.BrightCyan(arg)
}

func BrightWhite(arg any) aurora.Value {
	return c.BrightWhite(arg)
}

func Index(n aurora.ColorIndex, arg any) aurora.Value {
	return c.Index(n, arg)
}

func Gray(n aurora.GrayIndex, arg any) aurora.Value {
	return c.Gray(n, arg)
}

func BgBlack(arg any) aurora.Value {
	return c.BgBlack(arg)
}

func BgRed(arg any) aurora.Value {
	return c.BgRed(arg)
}

func BgGreen(arg any) aurora.Value {
	return c.BgGreen(arg)
}

func BgYellow(arg any) aurora.Value {
	return c.BgYellow(arg)
}

func BgBlue(arg any) aurora.Value {
	return c.BgBlue(arg)
}

func BgMagenta(arg any) aurora.Value {
	return c.BgMagenta(arg)
}

func BgCyan(arg any) aurora.Value {
	return c.BgCyan(arg)
}

func BgWhite(arg any) aurora.Value {
	return c.BgWhite(arg)
}

func BgBrightBlack(arg any) aurora.Value {
	return c.BgBrightBlack(arg)
}

func BgBrightRed(arg any) aurora.Value {
	return c.BgBrightRed(arg)
}

func BgBrightGreen(arg any) aurora.Value {
	return c.BgBrightGreen(arg)
}

func BgBrightYellow(arg any) aurora.Value {
	return c.BgBrightYellow(arg)
}

func BgBrightBlue(arg any) aurora.Value {
	return c.BgBrightBlue(arg)
}

func BgBrightMagenta(arg any) aurora.Value {
	return c.BgBrightMagenta(arg)
}

func BgBrightCyan(arg any) aurora.Value {
	return c.BgBrightCyan(arg)
}

func BgBrightWhite(arg any) aurora.Value {
	return c.BgBrightWhite(arg)
}

func BgIndex(n aurora.ColorIndex, arg any) aurora.Value {
	return c.BgIndex(n, arg)
}

func BgGray(n aurora.GrayIndex, arg any) aurora.Value {
	return c.BgGray(n, arg)
}

func Colorize(arg any, color aurora.Color) aurora.Value {
	return c.Colorize(arg, color)
}

func Hyperlink(arg any, target string, params ...aurora.HyperlinkParam) aurora.Value {
	return c.Hyperlink(arg, target, params...)
}

func HyperlinkTarget(arg any) (target string) {
	return c.HyperlinkTarget(arg)
}

func HyperlinkParams(arg any) (params []aurora.HyperlinkParam) {
	return c.HyperlinkParams(arg)
}

func Sprintf(format any, args ...any) string {
	return c.Sprintf(format, args...)
}
