package utils

import (
	"fmt"
	"github.com/fatih/color"
)

var uisuccess = color.New(color.FgGreen).SprintFunc()
var uiwarning = color.New(color.FgYellow).SprintFunc()
var uiinfo = color.New(color.FgHiBlue).SprintFunc()
var uisection = color.New(color.Bold, color.FgGreen).SprintFunc()
var uierror = color.New(color.FgHiRed).SprintFunc()

type Ui struct {
}

func (ui Ui) UiSuccess(message string) {
	fmt.Println(uisuccess(message))
}

func (ui Ui) UiWarning(message string) {
	fmt.Println(uiwarning(message))
}

func (ui Ui) UiInfo(message string) {
	fmt.Println(uiinfo(message))
}

func (ui Ui) UiError(err error) {
	fmt.Println(uierror(err))
}

func (ui Ui) UiSection(message string, arg any) {
	fmt.Printf("    %s: %s\n", uisection(message), uiinfo(arg))
}
