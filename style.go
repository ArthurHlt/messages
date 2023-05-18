package messages

import "github.com/logrusorgru/aurora/v4"

func Reset(arg interface{}) aurora.Value {
	return c.Reset(arg)
}

func Clear(arg interface{}) aurora.Value {
	return c.Clear(arg)
}

func Bold(arg interface{}) aurora.Value {
	return c.Bold(arg)
}

func Faint(arg interface{}) aurora.Value {
	return c.Faint(arg)
}

func DoublyUnderline(arg interface{}) aurora.Value {
	return c.DoublyUnderline(arg)
}

func Fraktur(arg interface{}) aurora.Value {
	return c.Fraktur(arg)
}

func Italic(arg interface{}) aurora.Value {
	return c.Italic(arg)
}

func Underline(arg interface{}) aurora.Value {
	return c.Underline(arg)
}

func SlowBlink(arg interface{}) aurora.Value {
	return c.SlowBlink(arg)
}

func RapidBlink(arg interface{}) aurora.Value {
	return c.RapidBlink(arg)
}

func Blink(arg interface{}) aurora.Value {
	return c.Blink(arg)
}

func Reverse(arg interface{}) aurora.Value {
	return c.Reverse(arg)
}

func Inverse(arg interface{}) aurora.Value {
	return c.Inverse(arg)
}

func Conceal(arg interface{}) aurora.Value {
	return c.Conceal(arg)
}

func Hidden(arg interface{}) aurora.Value {
	return c.Hidden(arg)
}

func CrossedOut(arg interface{}) aurora.Value {
	return c.CrossedOut(arg)
}

func StrikeThrough(arg interface{}) aurora.Value {
	return c.StrikeThrough(arg)
}

func Framed(arg interface{}) aurora.Value {
	return c.Framed(arg)
}

func Encircled(arg interface{}) aurora.Value {
	return c.Encircled(arg)
}

func Overlined(arg interface{}) aurora.Value {
	return c.Overlined(arg)
}

func Black(arg interface{}) aurora.Value {
	return c.Black(arg)
}

func Red(arg interface{}) aurora.Value {
	return c.Red(arg)
}

func Green(arg interface{}) aurora.Value {
	return c.Green(arg)
}

func Yellow(arg interface{}) aurora.Value {
	return c.Yellow(arg)
}

func Blue(arg interface{}) aurora.Value {
	return c.Blue(arg)
}

func Magenta(arg interface{}) aurora.Value {
	return c.Magenta(arg)
}

func Cyan(arg interface{}) aurora.Value {
	return c.Cyan(arg)
}

func White(arg interface{}) aurora.Value {
	return c.White(arg)
}

func BrightBlack(arg interface{}) aurora.Value {
	return c.BrightBlack(arg)
}

func BrightRed(arg interface{}) aurora.Value {
	return c.BrightRed(arg)
}

func BrightGreen(arg interface{}) aurora.Value {
	return c.BrightGreen(arg)
}

func BrightYellow(arg interface{}) aurora.Value {
	return c.BrightYellow(arg)
}

func BrightBlue(arg interface{}) aurora.Value {
	return c.BrightBlue(arg)
}

func BrightMagenta(arg interface{}) aurora.Value {
	return c.BrightMagenta(arg)
}

func BrightCyan(arg interface{}) aurora.Value {
	return c.BrightCyan(arg)
}

func BrightWhite(arg interface{}) aurora.Value {
	return c.BrightWhite(arg)
}

func Index(n aurora.ColorIndex, arg interface{}) aurora.Value {
	return c.Index(n, arg)
}

func Gray(n aurora.GrayIndex, arg interface{}) aurora.Value {
	return c.Gray(n, arg)
}

func BgBlack(arg interface{}) aurora.Value {
	return c.BgBlack(arg)
}

func BgRed(arg interface{}) aurora.Value {
	return c.BgRed(arg)
}

func BgGreen(arg interface{}) aurora.Value {
	return c.BgGreen(arg)
}

func BgYellow(arg interface{}) aurora.Value {
	return c.BgYellow(arg)
}

func BgBlue(arg interface{}) aurora.Value {
	return c.BgBlue(arg)
}

func BgMagenta(arg interface{}) aurora.Value {
	return c.BgMagenta(arg)
}

func BgCyan(arg interface{}) aurora.Value {
	return c.BgCyan(arg)
}

func BgWhite(arg interface{}) aurora.Value {
	return c.BgWhite(arg)
}

func BgBrightBlack(arg interface{}) aurora.Value {
	return c.BgBrightBlack(arg)
}

func BgBrightRed(arg interface{}) aurora.Value {
	return c.BgBrightRed(arg)
}

func BgBrightGreen(arg interface{}) aurora.Value {
	return c.BgBrightGreen(arg)
}

func BgBrightYellow(arg interface{}) aurora.Value {
	return c.BgBrightYellow(arg)
}

func BgBrightBlue(arg interface{}) aurora.Value {
	return c.BgBrightBlue(arg)
}

func BgBrightMagenta(arg interface{}) aurora.Value {
	return c.BgBrightMagenta(arg)
}

func BgBrightCyan(arg interface{}) aurora.Value {
	return c.BgBrightCyan(arg)
}

func BgBrightWhite(arg interface{}) aurora.Value {
	return c.BgBrightWhite(arg)
}

func BgIndex(n aurora.ColorIndex, arg interface{}) aurora.Value {
	return c.BgIndex(n, arg)
}

func BgGray(n aurora.GrayIndex, arg interface{}) aurora.Value {
	return c.BgGray(n, arg)
}

func Colorize(arg interface{}, color aurora.Color) aurora.Value {
	return c.Colorize(arg, color)
}

func Hyperlink(arg interface{}, target string, params ...aurora.HyperlinkParam) aurora.Value {
	return c.Hyperlink(arg, target, params...)
}

func HyperlinkTarget(arg interface{}) (target string) {
	return c.HyperlinkTarget(arg)
}

func HyperlinkParams(arg interface{}) (params []aurora.HyperlinkParam) {
	return c.HyperlinkParams(arg)
}

func Sprintf(format interface{}, args ...interface{}) string {
	return c.Sprintf(format, args...)
}
