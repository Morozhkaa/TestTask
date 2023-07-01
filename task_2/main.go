package main

import (
	"client/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"

	// create client, execute request
	client := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// read the body of the response
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// parse json into an array of structures
	data := []models.Сryptocurrency{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	// print the rates of certain or all cryptocurrencies (depending on the presence of arguments)
	// in the format: "Name (symbol) - current_price $"
	switch {
	case len(os.Args) > 1:
		m := make(map[string]struct{})
		for i := 1; i < len(os.Args); i++ {
			m[os.Args[i]] = struct{}{}
		}
		for _, v := range data {
			if _, ok := m[v.Name]; ok {
				fmt.Printf("%s (%s) - %.5f$\n", v.Name, v.Symbol, v.CurPrice)
			}
		}
	default:
		for _, v := range data {
			fmt.Printf("%s (%s) - %.5f$\n", v.Name, v.Symbol, v.CurPrice)
		}
	}
}
