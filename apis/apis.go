package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 30 * time.Second,
}

func fetcher(url string) (map[string]interface{}, error) {
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("HTTP error occurred '%s': %v\n", url, err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP error occurred '%s': status code %d\n", url, resp.StatusCode)
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body '%s': %v\n", url, err)
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("Error decoding JSON '%s': %v\n", url, err)
		return nil, err
	}

	return result, nil
}
