package serialization_bench

import (
	"testing"

	"serialization_bench/resources"
)

var TestResource = map[string]any{}

func TestMain(m *testing.M) {
	TestResource = map[string]any{
		"reddit":      resources.LoadResourceReddit(),
		"novel_prize": resources.LoadResourceNovelPrize(),
		"la_crime":    resources.LoadResourceLACrime(),
	}
	m.Run()
}
