package input

import (
	"fmt"
	"github.com/shrekie/comweb/command"
)

// Envoy executes command and
// formats data before represented by portrayer
type Envoy interface {
	Serve(Potray, error) (string, error)
	set(command.Query)
}

// Line is shown
type Line struct {
	query command.Query
}

// Serve as line
func (r *Line) Serve(p Potray, err error) (string, error) {
	if err != nil {
		return p.express("", err)
	}
	result, err := command.Execute(r.query)
	return p.express(fmt.Sprintf("%s \n", result), err)
}

func (r *Line) set(q command.Query) {
	r.query = q
}
