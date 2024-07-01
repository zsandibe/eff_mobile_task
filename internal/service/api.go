package service

import (
	"fmt"
	"io"
	"net/http"

	"github.com/zsandibe/eff_mobile_task/config"
)

func getInfoFromApi(regNum string, cfg config.Config) error {

	return nil
}

func getResponseBody(url, name string) ([]byte, error) {
	fullUrl := fmt.Sprintf("", url, name)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to requesting: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	return body, nil
}
