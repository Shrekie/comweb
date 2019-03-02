package input

import (
	"errors"
	"fmt"
	"github.com/shrekie/comweb/command"
	"os"
)

// Stream is a way to interpret the caller of program
type Stream interface {
	process(Argumenter) (Argumenter, error)
}

// OsStream takes data from the os argument invoker
type OsStream struct{}

func (s *OsStream) process(a Argumenter) (Argumenter, error) {
	err := a.new(os.Args)
	return a, err
}

// Argumenter is a method of constructing command queries
type Argumenter interface {
	query() *command.Query
	new([]string) error
}

// Splitter takes arguments as string array
// and splits them to three parts (name, arg, site)
type Splitter struct {
	arguments []string
}

// @todo: need to make a multiple splitter for arguments[2]
func (s *Splitter) query() *command.Query {
	return &command.Query{Name: s.arguments[1],
		Arg: s.arguments[2], Site: s.arguments[3]}
}

func (s *Splitter) new(a []string) error {
	if len(a) < 4 {
		return errors.New("insufficent length")
	}
	s.arguments = a
	return nil
}

// Resulter formats data before represented by portrayer
type Resulter interface {
	Serve(Potrayer, error) (string, error)
	new(*command.Query)
}

// Line is shown
type Line struct {
	Query *command.Query
}

// Serve as line
func (r *Line) Serve(p Potrayer, err error) (string, error) {
	if err != nil {
		return p.express("", err)
	}
	result, err := r.Query.Execute()
	return p.express(result, err)
}

func (r *Line) new(q *command.Query) {
	r.Query = q
}

// Potrayer is a way to show formatted string
type Potrayer interface {
	express(string, error) (string, error)
}

// TermPrinter print on terminal
type TermPrinter struct{}

func (t *TermPrinter) express(r string, err error) (string, error) {
	if err != nil {
		fmt.Printf("%s \n", err)
		return r, err
	}
	fmt.Printf("%s \n", r)
	return r, err
}

// New argumented stream to potray resulter
func New(s Stream, a Argumenter, r Resulter) (Resulter, error) {
	args, err := s.process(a)
	if err != nil {
		return r, err
	}
	query := args.query()
	r.new(query)
	return r, err
}
