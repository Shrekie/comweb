package input

import "fmt"

// Potray is a way to show formatted string
type Potray interface {
	express(string, error) (string, error)
}

// TermPrinter print on terminal
type TermPrinter struct{}

func (t *TermPrinter) express(r string, err error) (string, error) {
	if err != nil {
		fmt.Printf("%s \n", err)
		return r, err
	}
	fmt.Print(r)
	return r, err
}
