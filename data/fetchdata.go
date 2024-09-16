package database

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data from URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	return nil
}
