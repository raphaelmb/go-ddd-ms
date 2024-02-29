package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	res, err := sampleRes()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/partnerships", func(w http.ResponseWriter, r *http.Request) {
		min := 1
		max := 10
		ran := rand.Intn(max - min + 1)

		if ran > 7 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	})

	log.Println("Running on port 3031")
	if err := http.ListenAndServe(":3031", nil); err != nil {
		log.Fatal(err)
	}
}

type Res struct {
	AvailableHotels []struct {
		Name               string `json:"name"`
		PriceInUSDPerNight int    `json:"priceInUSDPerNight"`
	} `json:"availableHotels"`
}

func sampleRes() ([]byte, error) {
	sampleRes := Res{AvailableHotels: []struct {
		Name               string `json:"name"`
		PriceInUSDPerNight int    `json:"priceInUSDPerNight"`
	}{
		{
			Name:               "some hotel",
			PriceInUSDPerNight: 300,
		},
		{
			Name:               "some other hotel",
			PriceInUSDPerNight: 30,
		},
		{
			Name:               "some third hotel",
			PriceInUSDPerNight: 90,
		},
		{
			Name:               "some fourth hotel",
			PriceInUSDPerNight: 80,
		},
	}}

	b, err := json.Marshal(sampleRes)
	if err != nil {
		return nil, err
	}

	return b, nil
}
