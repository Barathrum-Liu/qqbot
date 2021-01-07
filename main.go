package main

import (
	"fmt"

	"github.com/Barathrum-Liu/qqbot/steam"
)

const apiKey = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

func main() {
	steamIDs := []uint64{76561198073971421}
	players, err := steam.GetPlayerSummaries(apiKey, steamIDs)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(players)
}
