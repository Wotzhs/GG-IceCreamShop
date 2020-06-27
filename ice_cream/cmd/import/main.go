package main

import (
	"GG-IceCreamShop/ice_cream/internal/models"
	"GG-IceCreamShop/ice_cream/internal/services"
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

const defaultIceCreamJSONURL string = "https://gist.githubusercontent.com/penmanglewood/f264e8d926b4c4a9926aa1de8fdb509a/raw/992f3c8a519ecd3d947bc48627ffefcf947f80bd/icecream.json"

func main() {
	var URL string
	flag.StringVar(&URL, "url", "", "Ice Cream JSON URL")
	flag.Parse()

	if URL == "" {
		log.Printf("Flag -url value is not set, using default ice cream json data url: %v\n", defaultIceCreamJSONURL)
		URL = defaultIceCreamJSONURL
	}

	log.Println("Downloading ice cream json")

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalf("http.Get() err: %v", err)
	}

	defer resp.Body.Close()

	log.Println("Downloaded ice cream json")

	var iceCreams []models.IceCream
	if err := json.NewDecoder(resp.Body).Decode(&iceCreams); err != nil {
		log.Fatalf("json.Decode() err: %v", err)
	}

	log.Println("Importing ice cream json")

	totalCount, err := services.IceCream.Import(iceCreams)
	if err != nil {
		log.Fatalf("services.IceCream.Import err: %v", err)
	}

	log.Printf("Imported %d ice creams\n", totalCount)
	log.Println("Import completed")
}
