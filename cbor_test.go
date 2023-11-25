package serialization_bench

import (
	"testing"

	fxamackerv1 "github.com/fxamacker/cbor"
	fxamackerv2 "github.com/fxamacker/cbor/v2"
)

func BenchmarkCBOR(b *testing.B) {
	b.Run("fxamacker/cbor/v1", func(b *testing.B) {
		for k, v := range TestResource {
			var serializedSize int
			b.Run(k, func(b *testing.B) {
				var result []byte
				var err error
				for i := 0; i < b.N; i++ {
					result, err = fxamackerv1.Marshal(v, fxamackerv1.CanonicalEncOptions())
					if err != nil {
						b.Fatal(err)
					}
					b.SetBytes(int64(len(result)))
					serializedSize = len(result)
				}
			})
			b.Logf("fxamacker/cbor/v1: %s: %s", k, formatByteSize(serializedSize))
		}
	})

	b.Run("fxamacker/cbor/v2", func(b *testing.B) {
		for k, v := range TestResource {
			var serializedSize int
			b.Run(k, func(b *testing.B) {
				var result []byte
				var err error
				for i := 0; i < b.N; i++ {
					result, err = fxamackerv2.Marshal(v)
					if err != nil {
						b.Fatal(err)
					}
					b.SetBytes(int64(len(result)))
					serializedSize = len(result)
				}
			})
			b.Logf("fxamacker/cbor/v2: %s: %s", k, formatByteSize(serializedSize))
		}
	})
}
