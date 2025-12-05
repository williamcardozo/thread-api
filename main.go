package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BrasilAPIResponse struct {
	CEP          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type ViaCepResponse struct {
	CEP          string `json:"cep"`
	State        string `json:"uf"`
	City         string `json:"localidade"`
	Neighborhood string `json:"bairro"`
	Street       string `json:"logradouro"`
}

type AddressResponse struct {
	Api      string
	Response interface{}
}

func fetchBrasilAPI(cep string, ch chan AddressResponse) {
	resp, err := http.Get(fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep))
	if err != nil {
		fmt.Printf("error fetching address: %v", err)
	}
	defer resp.Body.Close()

	var result BrasilAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("error decoding response: %v", err)
	}

	ch <- AddressResponse{
		"BrasilAPI",
		result,
	}
}

func fetchViaCep(cep string, ch chan AddressResponse) {
	resp, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		fmt.Printf("error fetching address: %v", err)
	}
	defer resp.Body.Close()

	var result ViaCepResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Printf("error decoding response: %v", err)
	}

	ch <- AddressResponse{
		"ViaCep",
		result,
	}
}

func formatAddressOutput(address interface{}) string {
	switch addr := address.(type) {
	case BrasilAPIResponse:
		return fmt.Sprintf("\n • Street: %s\n • Neighborhood: %s\n • City: %s\n • State: %s\n", addr.Street, addr.Neighborhood, addr.City, addr.State)
	case ViaCepResponse:
		return fmt.Sprintf("\n • Street: %s\n • Neighborhood: %s\n • City: %s\n • State: %s\n", addr.Street, addr.Neighborhood, addr.City, addr.State)
	}
	return ""
}

func main() {
	addressResponse := make(chan AddressResponse)
	cep := "01153000"

	go fetchViaCep(cep, addressResponse)
	go fetchBrasilAPI(cep, addressResponse)

	select {
	case rsp := <-addressResponse:
		fmt.Printf(" API: %s\n CEP: %s\n Address: %s", rsp.Api, cep, formatAddressOutput(rsp.Response))

	case <-time.After(1 * time.Second):
		fmt.Println("Timeout reached, exiting.")
		return
	}
}
