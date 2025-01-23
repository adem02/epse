package utils

type ProjectType string

const (
	GetLiteTemplatesPath  string      = "templates/lite"
	GetCleanTemplatesPath string      = "templates/clean"
	LiteProjectType       ProjectType = "lite"
	CleanProjectType      ProjectType = "clean"
	SrcPath               string      = "src/"
)
