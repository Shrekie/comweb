package commands

import (
	"errors"
	"github.com/shrekie/comweb/scraper"
	"strconv"
	"strings"
)

// Query to run a command
// command to match the query and
// site to target and optional argument
type Query struct {
	Name string
	Arg  string
	Site string
}

// Execute command from query
// finds matching name key in command map
// and executes the run function
func (q *Query) Execute() (string, error) {
	command, ok := commands[strings.ToLower(q.Name)]
	if ok {
		return command.run(q)
	}
	return "", errors.New("Command not found")
}

// Command that evaluates scraped web
type Command interface {
	run(*Query) (string, error)
	info() string
}

// each available command mapped
var commands = map[string]Command{
	"count": new(Count),
}

//
// COMMANDS
//

// Count arg sum occurence on web scrape
type Count struct{}

func (c *Count) run(q *Query) (string, error) {
	textBody, err := scraper.BodyText(q.Site)
	if err != nil {
		return "", err
	}
	result := strings.Count(textBody, q.Arg)
	return strconv.Itoa(result), nil
}

func (c *Count) info() string {
	return "sum occurence on web scrape"
}
