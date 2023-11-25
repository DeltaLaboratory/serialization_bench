package serialization_bench

import (
	"encoding/json"
	"testing"

	goccy "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
)

func BenchmarkJSON(b *testing.B) {
	b.Run("encoding/json", func(b *testing.B) {
		for k, v := range TestResource {
			var serializedSize int
			b.Run(k, func(b *testing.B) {
				var result []byte
				var err error
				for i := 0; i < b.N; i++ {
					result, err = json.Marshal(v)
					if err != nil {
						b.Fatal(err)
					}
					b.SetBytes(int64(len(result)))
					serializedSize = len(result)
				}
			})
			b.Logf("encoding/json: %s: %s", k, formatByteSize(serializedSize))
		}
	})

	b.Run("goccy/go-json", func(b *testing.B) {
		for k, v := range TestResource {
			var serializedSize int
			b.Run(k, func(b *testing.B) {
				var result []byte
				var err error
				for i := 0; i < b.N; i++ {
					result, err = goccy.Marshal(v)
					if err != nil {
						b.Fatal(err)
					}
					b.SetBytes(int64(len(result)))
					serializedSize = len(result)
				}
			})
			b.Logf("goccy/go-json: %s: %s", k, formatByteSize(serializedSize))
		}
	})

	b.Run("json-iterator/go", func(b *testing.B) {
		for k, v := range TestResource {
			var serializedSize int
			b.Run(k, func(b *testing.B) {
				var result []byte
				var err error
				for i := 0; i < b.N; i++ {
					result, err = jsoniter.Marshal(v)
					if err != nil {
						b.Fatal(err)
					}
					b.SetBytes(int64(len(result)))
					serializedSize = len(result)
				}
			})
			b.Logf("json-iterator/go: %s: %s", k, formatByteSize(serializedSize))
		}
	})
}
