package gomit

import ()

var (
	colorReset = "\033[0m"
    colorRed = "\033[31m"


	minCommitSize = 10
)

type Option struct {
	Type				[]string		`Available type of commit`
	MinimumSize			int				`Minimum size of commit`
	MaximumSize			int				`Maximum size of commit`
}

type Type struct {
	Name				string			`yaml:"name"`
}