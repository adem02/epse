package logutils

import (
	"fmt"

	"github.com/fatih/color"
)

var uisuccess = color.New(color.FgGreen).SprintFunc()
var uiwarning = color.New(color.FgYellow).SprintFunc()
var uiinfo = color.New(color.FgHiBlue).SprintFunc()
var uisection = color.New(color.Bold, color.FgGreen).SprintFunc()
var uierror = color.New(color.FgHiRed).SprintFunc()

type Logger struct {
}

func (logger Logger) Success(message string) {
	fmt.Println(uisuccess(message))
}

func (logger Logger) Warning(message string) {
	fmt.Println(uiwarning(message))
}

func (logger Logger) Info(message string) {
	fmt.Println(uiinfo(message))
}

func (logger Logger) Error(err error) {
	fmt.Println(uierror(err))
}

func (logger Logger) Section(message string, arg any) {
	fmt.Printf("    %s: %s\n", uisection(message), uiinfo(arg))
}
