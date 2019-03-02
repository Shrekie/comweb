package input

import (
	"fmt"
	"github.com/shrekie/comweb/command"
	"os"
)

// Stream is a way to interpret the caller of program
type Stream interface {
	process(Argumenter) Argumenter
}

// Argumenter is a method of constructing command queries
type Argumenter interface {
	query() *command.Query
	new([]string)
}

// Splitter takes arguments as string array
// and splits them to three parts (name, arg, site)
type Splitter struct {
	arguments []string
}

func (s *Splitter) query() *command.Query {
	return &command.Query{Name: s.arguments[1],
		Arg: s.arguments[2], Site: s.arguments[3]}
}

func (s *Splitter) new(a []string) {
	s.arguments = a
}

// OSStream takes data from the os argument invoker
type OSStream struct{}

func (s *OSStream) process(a Argumenter) Argumenter {
	a.new(os.Args)
	return a
}

// Resulter formats data before represented by portrayer
type Resulter interface {
	Serve(Potrayer)
	new(*command.Query)
}

// Line is shown
type Line struct {
	Query *command.Query
}

// Serve as line
func (r *Line) Serve(p Potrayer) {
	result, _ := r.Query.Execute()
	p.express(fmt.Sprintf("%s \n", result))
}

func (r *Line) new(q *command.Query) {
	r.Query = q
}

// Potrayer is a way to show formatted string
type Potrayer interface {
	express(string)
}

// TermPrinter print on terminal
type TermPrinter struct{}

func (t *TermPrinter) express(r string) {
	fmt.Print(r)
}

// Convey a stream in a splitting result
func Convey(s Stream, r Resulter) Resulter {
	query := s.process(new(Splitter)).query()
	r.new(query)
	return r
}
