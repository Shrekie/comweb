package main

import "github.com/shrekie/comweb/input"

func main() {
	args, err := input.New(&input.OsStream{}, &input.Splitter{}, &input.Line{})
	args.Serve(&input.TermPrinter{}, err)
}
