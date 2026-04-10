package typeutils

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
	LiteControllersPath   string         = "src/controllers"
	CleanControllersPath  string         = "src/adapters/controllers"
)

type TmplData struct {
	ProjectName     string
	Dependencies    string
	DevDependencies string
	ProjectType     string
	ControllersPath string
}

var ControllersPathMappedByProjectType = map[ProjectType]string{
	LiteProjectType:  "src/controllers",
	CleanProjectType: "src/adapters/controllers",
}

type RouteData struct {
	Domaine       string `json:"domaine"`
	RouteBasePath string `json:"routeBasePath"`
}

var ConfigFileBaseRoute = RouteData{
	Domaine:       "health",
	RouteBasePath: "/health",
}
