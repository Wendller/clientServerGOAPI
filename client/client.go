package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type CotationServerResponse struct {
	Bid string `json:"bid"`
}

var SERVER_LOCAL_URL string = "http://localhost:8080/cotacao"
var response CotationServerResponse

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", SERVER_LOCAL_URL, nil)
	if err != nil {
		log.Fatalf("HTTP create request to server failed: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatalf("HTTP request to server has reached timeout: %v", ctx.Err())
		}
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileRaw := fmt.Sprintf("Dólar: %v", response.Bid)

	_, err = file.WriteString(fileRaw)
	if err != nil {
		log.Fatal(err)
	}
}
