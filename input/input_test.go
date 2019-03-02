package input

import (
	"strconv"
	"testing"
)

var fakeInput = []string{"", "count", "google", "google.com"}

type fakeStream struct{}

func (s *fakeStream) process(a Argumenter) (Argumenter, error) {
	a.new(fakeInput)
	return a, nil
}

type fakePrinter struct{}

func (t *fakePrinter) express(r string, e error) (string, error) {
	return r, e
}

func TestLineOutput(t *testing.T) {
	args, err := New(&fakeStream{}, &Splitter{}, &Line{})
	result, err := args.Serve(&fakePrinter{}, err)
	if err != nil {
		t.Errorf("could not select and run command")
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
