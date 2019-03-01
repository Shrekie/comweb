package scraper

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// BodyText from url
func BodyText(url string) (string, error) {
	url = formatURL(url)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func formatURL(shortURL string) string {
	return fmt.Sprintf("http://%s", shortURL)
}
