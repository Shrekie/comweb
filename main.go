package main

import "github.com/shrekie/comweb/input"

func main() {
	input.Convey(new(input.OSStream), new(input.Line)).Serve(new(input.TermPrinter))
}
