package scraper

import (
	"strings"
	"testing"
)

func TestGetBodyText(t *testing.T) {
	site := "google.com"
	textBody, err := BodyText(site)
	if err != nil {
		t.Error("Error getting body text")
	}
	response := strings.Count(string(textBody), "google") > 0
	if !response {
		t.Error("Could not find 'google' on google.com body text")
	}
}
