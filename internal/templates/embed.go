package templates

import "embed"

//go:embed all:addcommand all:clean all:lite
var FS embed.FS
