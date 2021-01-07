package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type dota2player struct {
	AccountID  uint64 `json:"account_id"`
	PlayerSlot uint8  `json:"player_slot"`
	HeroID     uint8  `json:"hero_id"`
}

// Match dota2 match info
type Match struct {
	MatchID   uint64        `json:"match_id"`
	StartTime uint64        `json:"start_time"`
	LobbyType uint8         `json:"lobby_type"`
	Players   []dota2player `json:"players"`
}

type matchHistory struct {
	Result struct {
		Status       uint   `json:"status"`
		StatusDetail string `json:"statusDetail"`
		Matches      []Match
	}
}

// GetMatchHistory GetMatchHistory API
func GetMatchHistory(apiKey string, Dota2ID uint64, num uint64) ([]Match, error) {
	method := "/IDOTA2Match_570/GetMatchHistory/v001/"
	gmhurl := BaseURL + method
	strID := strconv.FormatUint(Dota2ID, 10)
	strNum := strconv.FormatUint(num, 10)
	vals := url.Values{}
	vals.Add("key", apiKey)
	vals.Add("account_id", strID)
	vals.Add("matches_requested", strNum)
	if vals != nil {
		gmhurl += "?" + vals.Encode()
	}
	resp, err := http.Get(gmhurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("steamapi %s Status code %d", gmhurl, resp.StatusCode)
	}
	var mh matchHistory
	d := json.NewDecoder(resp.Body)
	d.Decode((&mh))
	return mh.Result.Matches, nil
}
