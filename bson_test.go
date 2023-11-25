package serialization_bench

import "testing"

import (
	mongobson "go.mongodb.org/mongo-driver/bson"
	mgobson "gopkg.in/mgo.v2/bson"
)

func BenchmarkBSON(b *testing.B) {
	b.Run("mongo-go-driver/bson", func(b *testing.B) {
		for k, v := range TestResource {
			var serializedSize int
			var result []byte
			var err error
			b.Run(k, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					result, err = mongobson.Marshal(v)
					if err != nil {
						b.Fatal(err)
					}
					b.SetBytes(int64(len(result)))
					serializedSize = len(result)
				}
			})
			b.Logf("mongo-go-driver/bson: %s: %s", k, formatByteSize(serializedSize))
		}
	})

	b.Run("mgo.v2/bson", func(b *testing.B) {
		for k, v := range TestResource {
			var serializedSize int
			b.Run(k, func(b *testing.B) {
				var result []byte
				var err error
				for i := 0; i < b.N; i++ {
					result, err = mgobson.Marshal(v)
					if err != nil {
						b.Fatal(err)
					}
					b.SetBytes(int64(len(result)))
					serializedSize = len(result)
				}
			})
			b.Logf("mgo.v2/bson: %s: %s", k, formatByteSize(serializedSize))
		}
	})
}
