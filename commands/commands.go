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
		return command.Run(q)
	}
	return "", errors.New("Command not found")
}

// Command that evaluates scraped web
type Command struct {
	Run func(*Query) (string, error)
}

// each available command mapped
var commands = map[string]Command{
	"count": Command{Run: Count},
	"has":   Command{Run: Has},
}

//
// TOP COMMAND FUNCTIONS
//

// Count arg occurence on web scrape
func Count(q *Query) (string, error) {
	textBody, err := scraper.BodyText(q.Site)
	if err != nil {
		return "", err
	}
	result := strings.Count(textBody, q.Arg)
	return strconv.Itoa(result), nil
}

// Has arg on scrape, breaks and returns if true
func Has(q *Query) (string, error) {
	return "70", nil
}
