package resources

import (
	_ "embed"
	"encoding/json"
)

//go:embed nobel_prize.json
var nobelPrizeJSON []byte

type NobelPrize struct {
	Prizes []struct {
		Year      string `json:"year"`
		Category  string `json:"category"`
		Laureates []struct {
			Id         string `json:"id"`
			Firstname  string `json:"firstname"`
			Surname    string `json:"surname,omitempty"`
			Motivation string `json:"motivation"`
			Share      string `json:"share"`
		} `json:"laureates,omitempty"`
		OverallMotivation string `json:"overallMotivation,omitempty"`
	} `json:"prizes"`
}

func LoadResourceNobelPrize() *NobelPrize {
	var result NobelPrize
	err := json.Unmarshal(nobelPrizeJSON, &result)
	if err != nil {
		panic(err)
	}
	return &result
}
