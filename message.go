// Package messages - You should use messages for user communication as it's write on stderr
// For exporting command or others (for machine and scripting), use fmt and O color variable to ensure to set color or not
package messages

import (
	"context"
	"fmt"
	"github.com/jwalton/go-supportscolor"
	"io"
	"os"
	"time"

	"github.com/logrusorgru/aurora/v4"
	"github.com/mattn/go-isatty"
	"github.com/theckman/yacspin"
)

var output = os.Stderr

var supportColor = supportscolor.SupportsColor(output.Fd())
var c *aurora.Aurora

var spinner Spinner
var forcedTty = false
var notDrawable = false
var verbose = false
var isTerminal = isatty.IsTerminal(output.Fd())
var ctx = context.Background()
var spinConfig = &yacspin.Config{
	Writer:          os.Stderr,
	Frequency:       200 * time.Millisecond,
	CharSet:         yacspin.CharSets[62],
	Suffix:          " ",
	SuffixAutoColon: true,
	Message:         "working",
}

func init() {
	c = aurora.New(aurora.WithColors(supportColor.SupportsColor))
	if !isTerminal {
		spinner = &NullSpinner{}
		return
	}
	var err error
	spinner, err = yacspin.New(*spinConfig)
	if err != nil {
		panic(err)
	}
}

func Output() io.Writer {
	return output
}

func Isatty() bool {
	return isTerminal
}

// IsContextValue returns true if the key is set in the context and the value is not nil
func IsContextValue(key any) bool {
	val := ctx.Value(key)
	if val == nil {
		return false
	}
	ok, isBool := val.(bool)
	if !isBool {
		return ok
	}
	return true
}

func GetContextValue(key any) any {
	return ctx.Value(key)
}

func SetContextValue(key, value any) {
	ctx = context.WithValue(ctx, key, value)
}

func ForceTty() {
	supportscolor.IsTTYOption(true)
	c = aurora.New(aurora.WithColors(true))
	forcedTty = true
	isTerminal = true
	var err error
	spinConfig.TerminalMode = yacspin.ForceTTYMode
	spinner, err = yacspin.New(*spinConfig)
	if err != nil {
		panic(err)
	}
}

func SetColor(color bool) {
	c = aurora.New(aurora.WithColors(color))
}

func UseStdout() {
	output = os.Stdout
	supportColor = supportscolor.SupportsColor(output.Fd())
	if !forcedTty {
		c = aurora.New(aurora.WithColors(supportColor.SupportsColor))
		isTerminal = isatty.IsTerminal(output.Fd())
	}
}

func UseStderr() {
	output = os.Stderr
	supportColor = supportscolor.SupportsColor(output.Fd())
	if !forcedTty {
		c = aurora.New(aurora.WithColors(supportColor.SupportsColor))
		isTerminal = isatty.IsTerminal(output.Fd())
	}
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(output, a...)
}

func Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(output, a...)
}

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(output, format, a...)
}

func Printfln(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(output, format+"\n", a...)
}

func Error(str string) {
	Printfln("%s %s", c.Red("Error"), str)
}

func Errorf(format string, a ...interface{}) {
	Printf("%s ", c.Red("Error"))
	Printfln(format, a...)
}

func Debug(str string) {
	if !verbose {
		return
	}
	Printfln("%s %s", c.Gray(18, "Debug"), str)
}

func Debugf(format string, a ...interface{}) {
	if !verbose {
		return
	}
	Printf("%s ", c.Gray(18, "Debug"))
	Printfln(format, a...)
}

func Fatal(str string) {
	Printfln("%s %s", c.Red("Error"), str)
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	Printf("%s ", c.Red("Error"))
	Printfln(format, a...)
	os.Exit(1)
}

func Warning(str string) {
	Printfln("%s %s", c.Yellow("Warning"), str)
}

func Warningf(format string, a ...interface{}) {
	Printf("%s ", c.Yellow("Warning"))
	Printfln(format, a...)
}

func Success(str string) {
	Printfln("%s %s", c.Green("Success"), str)
}

func Successf(format string, a ...interface{}) {
	Printf("%s ", c.Green("Success"))
	Printfln(format, a...)
}

func Info(str string) {
	Printfln("%s %s", c.Cyan("Info"), str)
}

func Infof(format string, a ...interface{}) {
	Printf("%s ", c.Cyan("Info"))
	Printfln(format, a...)
}

func Tips(str string) {
	Printfln("%s %s", c.Cyan("Tips"), str)
}

func Tipsf(format string, a ...interface{}) {
	Printf("%s ", c.Cyan("Tips"))
	Printfln(format, a...)
}

func GetSpinner(forceNull ...bool) Spinner {
	if len(forceNull) > 0 && forceNull[0] {
		return &NullSpinner{}
	}
	return spinner
}

func SetSpinner(s Spinner) {
	spinner = s
}

func SetVerbose(v bool) {
	verbose = v
}
