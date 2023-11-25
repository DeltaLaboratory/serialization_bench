package resources

import (
	_ "embed"
	"encoding/json"
)

//go:embed novel_prize.json
var novelPrizeJSON []byte

type NovelPrize struct {
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

func LoadResourceNovelPrize() *NovelPrize {
	var result NovelPrize
	err := json.Unmarshal(novelPrizeJSON, &result)
	if err != nil {
		panic(err)
	}
	return &result
}
