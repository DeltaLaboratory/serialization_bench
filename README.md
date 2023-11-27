# Go Serialization Benchmark
## to run the benchmark
```bash
sh resources/recv_files.sh
go test -bench=/ -benchmem
```
## result
goos: windows\
goarch: amd64\
pkg: serialization_bench\
cpu: 13th Gen Intel(R) Core(TM) i7-1365U\
32GiB RAM / DDR5 6000MHz


| Benchmark                                                  | Iterations | Time per Iteration (ns/op) | MB/s    | Memory (B/op) | Memory Allocations (allocs/op) |
|------------------------------------------------------------|------------|----------------------------|---------|---------------|--------------------------------|
| BenchmarkBSON/mongo-go-driver/bson/reddit-12               | 6404       | 225152                     | 567.80  | 131072        | 1                              |
| BenchmarkBSON/mongo-go-driver/bson/nobel_prize-12          | 2788       | 422061                     | 615.04  | 262665        | 1                              |
| BenchmarkBSON/mongo-go-driver/bson/la_crime-12             | 1          | 2036393900                 | 238.60  | 3907352032    | 756                            |
| BenchmarkBSON/mgo.v2/bson/reddit-12                        | 2817       | 377947                     | 338.91  | 879768        | 4636                           |
| BenchmarkBSON/mgo.v2/bson/nobel_prize-12                   | 1396       | 909106                     | 285.81  | 1864144       | 15697                          |
| BenchmarkBSON/mgo.v2/bson/la_crime-12                      | 1          | 1660805800                 | 292.56  | 3115667272    | 5659841                        |
| BenchmarkCBOR/fxamacker/cbor/v1/reddit-12                  | 15442      | 84748                      | 1414.86 | 122931        | 1                              |
| BenchmarkCBOR/fxamacker/cbor/v1/nobel_prize-12             | 4581       | 232222                     | 848.91  | 204914        | 1                              |
| BenchmarkCBOR/fxamacker/cbor/v1/la_crime-12                | 2          | 947324000                  | 299.86  | 1357546176    | 20                             |
| BenchmarkCBOR/fxamacker/cbor/v2/reddit-12                  | 16172      | 80839                      | 1483.27 | 122928        | 1                              |
| BenchmarkCBOR/fxamacker/cbor/v2/nobel_prize-12             | 5655       | 196829                     | 1001.56 | 204893        | 1                              |
| BenchmarkCBOR/fxamacker/cbor/v2/la_crime-12                | 2          | 807073850                  | 358.24  | 825999204     | 13                             |
| BenchmarkCBOR/brianolson/cbor_go/la_crime-12               | 2          | 684547250                  | 422.36  | 402740616     | 6925                           |
| BenchmarkCBOR/brianolson/cbor_go/reddit-12                 | 1172       | 959486                     | 125.87  | 184991        | 14909                          |
| BenchmarkCBOR/brianolson/cbor_go/nobel_prize-12            | 615        | 1863240                    | 112.49  | 444436        | 30725                          |
| BenchmarkGOB/reddit-12                                     | 10000      | 121022                     | 844.18  | 353030        | 353                            |
| BenchmarkGOB/nobel_prize-12                                | 5900       | 205916                     | 726.95  | 686778        | 59                             |
| BenchmarkGOB/la_crime-12                                   | 1          | 4426046800                 | 117.04  | 3858184344    | 6303010                        |
| BenchmarkJSON/encoding/json/reddit-12                      | 6934       | 156260                     | 947.92  | 155736        | 1                              |
| BenchmarkJSON/encoding/json/nobel_prize-12                 | 3978       | 279944                     | 812.23  | 229508        | 1                              |
| BenchmarkJSON/encoding/json/la_crime-12                    | 1          | 1167538400                 | 302.39  | 1426807656    | 62                             |
| BenchmarkJSON/goccy/go-json/reddit-12                      | 10000      | 111832                     | 1324.51 | 155718        | 1                              |
| BenchmarkJSON/goccy/go-json/nobel_prize-12                 | 6308       | 183668                     | 1237.99 | 229564        | 1                              |
| BenchmarkJSON/goccy/go-json/la_crime-12                    | 1          | 1257940400                 | 280.66  | 3686410992    | 3592                           |
| BenchmarkJSON/json-iterator/go/reddit-12                   | 6320       | 169249                     | 875.17  | 155764        | 2                              |
| BenchmarkJSON/json-iterator/go/nobel_prize-12              | 3979       | 291109                     | 781.08  | 229682        | 2                              |
| BenchmarkJSON/json-iterator/go/la_crime-12                 | 1          | 1073675000                 | 328.83  | 2233012168    | 5075                           |
| BenchmarkMSGPACK/vmihailenco/msgpack/v5/reddit-12          | 10000      | 129238                     | 908.15  | 262728        | 38                             |
| BenchmarkMSGPACK/vmihailenco/msgpack/v5/nobel_prize-12     | 3698       | 351824                     | 595.69  | 524272        | 14                             |
| BenchmarkMSGPACK/vmihailenco/msgpack/v5/la_crime-12        | 1          | 1268595400                 | 228.21  | 805358824     | 856                            |
| BenchmarkMSGPACK/shamaton/msgpack/v2/reddit-12             | 10402      | 126636                     | 926.81  | 137688        | 749                            |
| BenchmarkMSGPACK/shamaton/msgpack/v2/nobel_prize-12        | 4034       | 295974                     | 708.10  | 267776        | 3112                           |
| BenchmarkMSGPACK/shamaton/msgpack/v2/la_crime-12           | 2          | 821951000                  | 352.22  | 330005832     | 1687224                        |
| BenchmarkMSGPACK/hashicorp/go-msgpack/codec/reddit-12      | 6927       | 166453                     | 707.20  | 156637        | 796                            |
| BenchmarkMSGPACK/hashicorp/go-msgpack/codec/nobel_prize-12 | 2866       | 389409                     | 540.93  | 316318        | 3360                           |
| BenchmarkMSGPACK/hashicorp/go-msgpack/codec/la_crime-12    | 3          | 437928067                  | 666.32  | 378232368     | 844742                         |