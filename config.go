package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	APIKey   string            `json:"apiKey"`
	SteamIDs map[string]uint64 `json:"steamIDs"`
}

// Configuration apikey 和steam id列表
var Configuration configuration

func init() {
	file, err := os.Open("conf.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	d := json.NewDecoder(file)
	err = d.Decode(&Configuration)
	if err != nil {
		panic(err)
	}
	fmt.Printf("configuration loaded.\n")
}
