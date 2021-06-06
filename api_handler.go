package deepl

import (
	"io"
	"io/ioutil"
	"net/http"
)

const contentType string = "application/x-www-form-urlencoded"

// POST is creating POST request to the endpoint with body.
func POST(endpoint string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", contentType)
	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
