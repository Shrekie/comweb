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

func TestCountLine(t *testing.T) {
	args, err := New(&fakeStream{}, &Splitter{}, &Line{})
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
