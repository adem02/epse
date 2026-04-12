package common

import (
	"fmt"
	"strings"

	"github.com/adem02/epse/internal/utils/logutils"
)

func AddImportStatementInFile(fileContent *string, importStatement string) {
	importString := `import { Router } from 'express';`
	if !strings.Contains(*fileContent, importStatement) {
		*fileContent = strings.Replace(*fileContent, importString, importString+"\n"+importStatement, 1)
	}
}

func InsertContentBeforeStatementInFile(fileContent *string, statement, newContent string) {
	insertPos := strings.LastIndex(*fileContent, statement)
	if insertPos == -1 {
		logutils.Logger{}.Error(fmt.Errorf("unable to find statement %s in file", statement))
		return
	}

	prevContent := strings.TrimSpace((*fileContent)[:insertPos])
	*fileContent = prevContent + "\n" + newContent + "\n\n" + (*fileContent)[insertPos:]
}
