package argument

import (
	`os`
)

type ArgumentSet struct {
	Action
	actions map[string]*Action
	prog    string
	args    []string
}

func NewArgumentSet(
	prog string,
	args []string,
) *ArgumentSet {

	if args == nil {
		args = os.Args[1:]
	}

	if len(prog) < 1 {
		prog = os.Args[0]
	}

	return &ArgumentSet{
		prog: prog,
		args: args,
	}
}

// vim:ts=4 sw=4 noet:
