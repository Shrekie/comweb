package command

import (
	"strconv"
	"testing"
)

func TestCountCommand(t *testing.T) {
	query := new(Site)
	query.set([]string{"", "count", "google", "google.com"})
	executed, err := Execute(query)
	if err != nil {
		t.Error("could not execute query")
	}
	response, err := strconv.Atoi(executed)
	if err != nil {
		t.Errorf("could not convert %s to integer", executed)
	}
	expectedMax, expectedMin := 100, 50
	if response < expectedMin {
		t.Errorf("%d is less than %d", response, expectedMin)
	}
	if response > expectedMax {
		t.Errorf("%d is more than %d", response, expectedMax)
	}
}
