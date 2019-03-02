package main

import "github.com/shrekie/comweb/input"

func main() {

	system := &input.Input{
		Channel:     &input.OsChannel{},
		Destructure: &input.Splitter{},
		Envoy:       &input.Line{}}

	args, err := system.Transmit()

	args.Serve(&input.TermPrinter{}, err)

}
