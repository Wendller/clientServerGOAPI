package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

func GetUSDToBRLCotation() (*USDBRL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", USD_TO_BRL_URL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request with context error: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatal("HTTP request cotation timeout reached")
			return nil, ctx.Err()
		}

		return nil, fmt.Errorf("HTTP request cotation error: %v", err)
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

	return &response.UsdBrl, nil
}
