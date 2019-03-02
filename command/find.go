package command

import (
	"errors"
	"strings"
)

// each available command
var commands = map[string]Command{
	"count": new(Count),
	"has":   new(Has),
	"info":  new(Info),
}

// Pair a set of arguments to query type
func Pair(v []string) (Query, error) {
	var query Query
	var err error

	if len(v) != 3 && len(v) != 4 {
		query, err = new(Site), errors.New("invalid arguments")
		return query, err
	}

	if len(v) == 4 {
		query, err = new(Site), nil
	}
	if len(v) == 3 {
		query, err = new(Help), nil
	}

	query.set(v)
	return query, err
}

// Execute command based on query
func Execute(q Query) (string, error) {
	command, ok := commands[strings.ToLower(q.target())]
	if ok {
		return command.run(q, q.get())
	}
	return "", errors.New("Command not found")
}
