package main

import (
	minCopiesPurchaseService "flexeraCodeTest/app/pkg/services/min_copies_purchase_service"
	"log"
	"time"
)

func main() {
	// calculate the minimum copies of the application with ID 374 a company must purchase from small csv sample
	now := time.Now()
	minCopiesPurchase, err := minCopiesPurchaseService.MinCopiesPurchase("sample-small.csv")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Minimum Copies Purchase: %d, Time Taken: %v\n", minCopiesPurchase, time.Since(now))

	// calculate the minimum copies of the application with ID 374 a company must purchase from large csv sample
	now = time.Now()
	minCopiesPurchase, err = minCopiesPurchaseService.MinCopiesPurchase("sample-large.csv")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Minimum Copies Purchase: %d, Time Taken: %v\n", minCopiesPurchase, time.Since(now))
}
