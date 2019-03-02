package input

import (
	"strconv"
	"testing"
)

var fakeInput = []string{"count", "google", "google.com"}

type fakeStream struct{}

func (s *fakeStream) process(a Argumenter) Argumenter {
	a.new(fakeInput)
	return a
}

type fakePrinter struct{}

func (t *fakePrinter) express(r string, e error) (string, error) {
	return r
}

func TestLineOutput(t *testing.T) {
	selected, err := Convey(new(fakeStream), new(Line)).Serve(new(fakePrinter))
	if err != nil {
		t.Errorf("could not select and run command")
	}
	count, err := strconv.Atoi(selected)
	if err != nil {
		t.Errorf("could not convert to integer")
	}
	response := count > 0
	if !response {
		t.Errorf("Could not find 'google' on google.com body text %d", count)
	}
}
