package command

import (
	"errors"
	"github.com/shrekie/comweb/scraper"
	"strconv"
	"strings"
)

// Command is a command callable on program
type Command interface {
	run(Query, map[string]string) (string, error)
	info() string
}

// Count arg sum occurence
type Count struct{}

func (c *Count) run(q Query, data map[string]string) (string, error) {
	textBody, err := scraper.BodyText(data["site"])
	if err != nil {
		return "", err
	}
	result := strings.Count(textBody, data["arg"])
	return strconv.Itoa(result), nil
}

func (c *Count) info() string {
	return "Count arg sum occurence"
}

// Has given argument on site
type Has struct{}

func (c *Has) run(q Query, data map[string]string) (string, error) {
	textBody, err := scraper.BodyText(data["site"])
	if err != nil {
		return "", err
	}
	result := strings.Count(textBody, data["arg"])
	return strconv.FormatBool(result > 0), nil
}

func (c *Has) info() string {
	return "Has given argument on site"
}

// Info about a given command
type Info struct{}

func (c *Info) run(q Query, data map[string]string) (string, error) {
	command, ok := commands[strings.ToLower(data["arg"])]
	if !ok {
		return "", errors.New("Command not found")
	}
	return command.info(), nil
}

func (c *Info) info() string {
	return "Gives you info about commands"
}
