package colors

import "github.com/fatih/color"

var Magenta = color.New(color.FgMagenta, color.Bold).SprintFunc()

var Green = color.New(color.FgHiGreen).SprintFunc()

var Fail = color.New(color.FgHiRed).SprintFunc()

var Info = color.New(color.FgCyan, color.Bold).SprintFunc()
