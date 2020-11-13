package nodeengine

const (
	Node = "node"
	Npm  = "npm"

	DepKey             = "dependency-sha"
	NvmrcSource        = ".nvmrc"
	BuildpackYMLSource = "buildpack.yml"
)

var (
	Priorities = map[string]int{
		"buildpack.yml": 3,
		"package.json":  2,
		".nvmrc":        1,
		"":              -1,
	}
)
