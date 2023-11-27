package serialization_bench

import (
	"bytes"
	"testing"
)

import (
	hashicorp "github.com/hashicorp/go-msgpack/codec"
	shamaton "github.com/shamaton/msgpack/v2"
	vmihailencov5 "github.com/vmihailenco/msgpack/v5"
)

func BenchmarkMSGPACK(b *testing.B) {
	b.Run("vmihailenco/msgpack/v5", func(b *testing.B) {
		for k, v := range TestResource {
			var serializedSize int
			b.Run(k, func(b *testing.B) {
				var result []byte
				var err error
				for i := 0; i < b.N; i++ {
					result, err = vmihailencov5.Marshal(v)
					if err != nil {
						b.Fatal(err)
					}
					b.SetBytes(int64(len(result)))
					serializedSize = len(result)
				}
			})
			b.Logf("vmihailenco/msgpack/v5: %s: %s", k, formatByteSize(serializedSize))
		}
	})

	b.Run("shamaton/msgpack/v2", func(b *testing.B) {
		for k, v := range TestResource {
			var serializedSize int
			b.Run(k, func(b *testing.B) {
				var result []byte
				var err error
				for i := 0; i < b.N; i++ {
					result, err = shamaton.Marshal(v)
					if err != nil {
						b.Fatal(err)
					}
					b.SetBytes(int64(len(result)))
					serializedSize = len(result)
				}
			})
			b.Logf("vmihailenco/msgpack/v5: %s: %s", k, formatByteSize(serializedSize))
		}
	})

	b.Run("hashicorp/go-msgpack/codec", func(b *testing.B) {
		for k, v := range TestResource {
			var serializedSize int
			b.Run(k, func(b *testing.B) {
				var err error
				buff := bytes.NewBuffer(nil)
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					buff.Reset()
					b.StartTimer()
					err = hashicorp.NewEncoder(buff, &hashicorp.MsgpackHandle{}).Encode(v)
					if err != nil {
						b.Fatal(err)
					}
					b.SetBytes(int64(buff.Len()))
					serializedSize = buff.Len()
				}
			})
			b.Logf("hashicorp/go-msgpack/codec: %s: %s", k, formatByteSize(serializedSize))
		}
	})

}
