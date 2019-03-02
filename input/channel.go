package input

import "os"

// Channel is a way to interpret the caller of program
type Channel interface {
	process(Destructure) Destructure
}

// OsChannel takes data from the os argument invoker
type OsChannel struct{}

func (s *OsChannel) process(d Destructure) Destructure {
	d.set(os.Args)
	return d
}
