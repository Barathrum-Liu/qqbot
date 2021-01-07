package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// BaseURL base api url
const BaseURL = "https://api.steampowered.com"

// Player player's info
type Player struct {
	SteamID       string `json:"steamid"`
	PersonaName   string `json:"personaname"`
	PersonaState  string `json:"personastate"`
	GameExtraInfo string `json:"gameextrainfo,omitempty"`
}

type playerSummaries struct {
	Response struct {
		Players []Player
	}
}

// GetPlayerSummaries GetPlayerSummaries API
func GetPlayerSummaries(apiKey string, steamIDs []uint64) ([]Player, error) {
	method := "/ISteamUser/GetPlayerSummaries/v0002/"
	gpsurl := BaseURL + method
	strIds := make([]string, len(steamIDs))
	for _, id := range steamIDs {
		strIds = append(strIds, strconv.FormatUint(id, 10))
	}
	vals := url.Values{}
	vals.Add("key", apiKey)
	vals.Add("steamids", strings.Join(strIds, ","))
	var playersummaries playerSummaries
	if vals != nil {
		gpsurl += "?" + vals.Encode()
	}
	resp, err := http.Get(gpsurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("steamapi %s Status code %d", gpsurl, resp.StatusCode)
	}

	d := json.NewDecoder(resp.Body)
	d.Decode((&playersummaries))
	return playersummaries.Response.Players, nil
}
