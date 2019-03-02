package input

import "github.com/shrekie/comweb/command"

// Destructure is a method of grouping command queries
type Destructure interface {
	query() (command.Query, error)
	set([]string)
}

// Splitter takes arguments as string array
// and splits them as array to the query generator
type Splitter struct {
	arguments []string
}

func (s *Splitter) query() (command.Query, error) {
	return command.Pair(s.arguments)
}

func (s *Splitter) set(a []string) {
	s.arguments = a
}
