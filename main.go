package main

import "github.com/shrekie/comweb/input"

func main() {
	result, err := input.New(&input.OsStream{}, &input.Splitter{}, &input.Line{})
	result.Serve(&input.TermPrinter{}, err)
}
