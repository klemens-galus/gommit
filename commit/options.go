package gomit

import ()

var (
	colorReset = "\033[0m"
    colorRed = "\033[31m"
    colorGreen = "\033[32m"
	colorMagenta = "\033[95m"
	colorCyan = "\033[96m"
	colorBlue = "\033[94m"


	minCommitSize = 10
	maxCommitSize = -1
	types = []string{
		"build",
		"chore",
		"ci",
		"docs",
		"feat",
		"fix",
		"perf",
		"refactor",
		"revert",
		"style",
		"test",
	}
)

func getOption() Option {
	o := Option{
		Type: types,
		MinimumSize: minCommitSize,
		MaximumSize: maxCommitSize,
	}
	return o
}

type Option struct {
	Type				[]string		`Available type of commit`
	MinimumSize			int				`Minimum size of commit`
	MaximumSize			int				`Maximum size of commit`
}
