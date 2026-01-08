# uuuidbench

A collection benchmarks for various uuid packages.

Read more about the [benchmark results here](https://lgsn.dev/uuidbench).

```
~/uuidbench go test ./... -bench=. -benchmem -benchtime=3s
goos: linux
goarch: amd64
pkg: github.com/lsl/uuidbench
cpu: 13th Gen Intel(R) Core(TM) i7-13800H
BenchmarkGofrsV7-20             34792820               102.9 ns/op            16 B/op          1 allocs/op
BenchmarkGoogleV7-20            26217514               130.9 ns/op            16 B/op          1 allocs/op
BenchmarkToolboxV7-20            8009846               446.0 ns/op           216 B/op          7 allocs/op
BenchmarkUUID6ProtoV7-20         2243642              1602 ns/op               8 B/op          1 allocs/op
BenchmarkCoolAJV7-20            41139253                86.47 ns/op           16 B/op          1 allocs/op
BenchmarkLSLV7-20               53252844                65.29 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/lsl/uuidbench        25.976s
```
