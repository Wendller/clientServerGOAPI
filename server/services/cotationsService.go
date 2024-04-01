package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type USDBRLResponse struct {
	UsdBrl USDBRL
}

type USDBRL struct {
	Bid string `json:"bid"`
}

var USD_TO_BRL_URL string = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
var response USDBRLResponse

func GetUSDToBRL() (*USDBRL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", USD_TO_BRL_URL, nil)
	if err != nil {
		return nil, fmt.Errorf("request with context error: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %v", err)
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("HTTP response reading error: %v", err)
	}

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, fmt.Errorf("response parse error: %v", err)
	}

	// jsonData, err := json.Marshal(response.UsdBrl)
	// if err != nil {
	// 	return nil, fmt.Errorf("response parse error: %v", err)
	// }

	return &response.UsdBrl, nil
}
