package main

import (
	"fmt"
	"log"

	hyperion "github.com/World-of-Cryptopups/eos-hyperion.go"
)

func main() {
	client := hyperion.New("https://testnet.waxsweden.org")

	health, err := client.GetHealth()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(health.Version)

	abi, err := client.HistoryGetAbiSnapshot(hyperion.GetAbiSnapshotParams{
		Contract: "shomaiiblend",
		Fetch:    true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(abi.ABI)
}
