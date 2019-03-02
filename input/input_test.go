package input

import (
	"github.com/shrekie/comweb/command"
	"strconv"
	"testing"
)

var fakeInput = []string{"", "count", "google", "google.com"}

type fakeChannel struct{}

func (s *fakeChannel) process(d Destructure) Destructure {
	d.set(fakeInput)
	return d
}

type fakePrinter struct{}

func (t *fakePrinter) express(r string, e error) (string, error) {
	return r, e
}

type fakeEnvoyer struct {
	query command.Query
}

func (f *fakeEnvoyer) Serve(p Potray, err error) (string, error) {
	if err != nil {
		return p.express("", err)
	}
	result, err := command.Execute(f.query)
	return p.express(result, err)
}

func (f *fakeEnvoyer) set(q command.Query) {
	f.query = q
}

func TestCountLine(t *testing.T) {
	system := &Input{
		Channel:     &fakeChannel{},
		Destructure: &Splitter{},
		Envoy:       &fakeEnvoyer{}}

	args, err := system.Transmit()

	if err != nil {
		t.Errorf("Could not construct query")
	}
	result, err := args.Serve(&fakePrinter{}, err)
	if err != nil {
		t.Errorf("Could not serve or execute query")
	}
	count, err := strconv.Atoi(result)
	if err != nil {
		t.Errorf("could not convert %s to integer", result)
	}
	response := count > 0
	if !response {
		t.Errorf("Could not find 'google' on google.com body text %d", count)
	}
}
