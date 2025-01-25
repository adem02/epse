package utils

type ProjectType string
type DependencyType string
type DependencyMap map[DependencyType][]string
type FormattedDependencyMap map[DependencyType]string

const (
	GetLiteTemplatesPath  string         = "templates/lite"
	GetCleanTemplatesPath string         = "templates/clean"
	LiteProjectType       ProjectType    = "lite"
	CleanProjectType      ProjectType    = "clean"
	SrcPath               string         = "src/"
	Dependencies          DependencyType = "dependencies"
	DevDependencies       DependencyType = "devDependencies"
)

type TmplData struct {
	ProjectName     string
	Dependencies    string
	DevDependencies string
}
