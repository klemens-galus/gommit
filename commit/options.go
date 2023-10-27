package gomit

import ()

type Option struct {
	Type				[]string		`Available type of commit`
	MinimumSize			int				`Minimum size of commit`
	MaximumSize			int				`Maximum size of commit`
}

type Type struct {
	Name				string			`yaml:"name"`
}