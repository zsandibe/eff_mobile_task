package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/zsandibe/eff_mobile_task/config"
	"github.com/zsandibe/eff_mobile_task/internal/domain"
)

func getInfoFromApi(passportSerie, passportNumber string, cfg config.Config) (domain.People, error) {
	body, err := getResponseBody(cfg.Api.Uri, passportSerie, passportNumber)
	if err != nil {
		return domain.People{}, err
	}

	var response domain.People

	if err := json.Unmarshal(body, &response); err != nil {
		return domain.People{}, fmt.Errorf("problems with unmarshalling response: %v", err)
	}
	return response, nil
}

func getResponseBody(url, serie, number string) ([]byte, error) {
	fullUrl := fmt.Sprintf("%s/info?passportSerie=%s?passportNumber=%s", url, serie, number)
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
