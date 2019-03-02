package input

import (
	"strconv"
	"testing"
)

var fakeSlice = [3]string{"count", "google", "google.com"}

type termArguments struct{}

func (i *termArguments) Slice() (string, string, string) {
	return fakeSlice[0], fakeSlice[1], fakeSlice[2]
}

func TestInput(t *testing.T) {
	input := new(termArguments)
	selected, err := Convey(input)
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
