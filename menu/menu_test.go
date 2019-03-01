package menu

import (
	"strconv"
	"testing"
)

var fakeMessageCounter = 0
var fakeMessageSelections = [3]string{"count", "google", "google.com"}

type fakeInputReader struct{}

func (i *fakeInputReader) Prompt(t string) (string, error) {
	fakeSelection := fakeMessageSelections[fakeMessageCounter]
	fakeMessageCounter++
	return fakeSelection, nil
}

func TestInput(t *testing.T) {
	input := new(fakeInputReader)
	selected, err := Select(input)
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
