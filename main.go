package main

import (
	"fmt"

	"github.com/Barathrum-Liu/qqbot/steam"
)

func main() {
	steamIDs := []uint64{}
	for _, id := range Configuration.SteamIDs {
		steamIDs = append(steamIDs, id+76561197960265728)
	}
	players, err := steam.GetPlayerSummaries(Configuration.APIKey, steamIDs)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(players)
	for _, id := range Configuration.SteamIDs {
		match, err := steam.GetMatchHistory(Configuration.APIKey, id, 1)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(match)
	}
}
