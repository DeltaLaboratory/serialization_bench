package serialization_bench

import (
	"bytes"
	"encoding/gob"
	"testing"
)

func BenchmarkGOB(b *testing.B) {
	for k, v := range TestResource {
		var serializedSize int
		b.Run(k, func(b *testing.B) {
			var err error
			buff := bytes.NewBuffer(nil)
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				buff.Reset()
				b.StartTimer()
				err = gob.NewEncoder(buff).Encode(v)
				if err != nil {
					b.Fatal(err)
				}
				b.SetBytes(int64(buff.Len()))
				serializedSize = buff.Len()
			}
		})
		b.Logf("encoding/gob: %s: %s", k, formatByteSize(serializedSize))
	}
}
