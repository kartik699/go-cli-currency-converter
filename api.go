package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type CurrencyResponse struct {
	Result 			 string  `json:"result"`
	ConversionResult float64 `json:"conversion_result"`
}

var base string = "https://v6.exchangerate-api.com/v6/"

func GetCurrencyRate(currencyFrom string, currencyTo string, amount float64) {
	err := godotenv.Load(".env.local")

	if err != nil {
		log.Fatal("env file not present")
	}

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		log.Fatal("API Key not found")
	}

	var url string = fmt.Sprintf(base + "%s/pair/%s/%s/%.2f", apiKey, currencyFrom, currencyTo, amount)

	resp, err := http.Get(url)
	if err != nil {
        log.Fatalf("Failed to make API request: %v", err)
    }
    defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API request failed with status: %s", resp.Status)
	}

	// decoding json
	var currencyData CurrencyResponse
	if err = json.NewDecoder(resp.Body).Decode(&currencyData); err != nil {
		log.Fatalf("Failed to decode JSON response: %v", err)
	}

	if currencyData.Result == "success" {
		fmt.Println("Converted Amount:", currencyData.ConversionResult, currencyTo)
    } else {
        log.Println("Failed to get a successful response from the API.")
    }
}