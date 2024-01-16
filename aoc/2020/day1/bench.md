### full n cube runtime

```
12:19 $ go test  ./day1/. -bench FpNn -v -benchtime=3s
goos: linux
goarch: amd64
pkg: github.com/Gnoale/adventofcode/day1
BenchmarkFpNn
BenchmarkFpNn-8           330440              9957 ns/op
PASS
ok      github.com/Gnoale/adventofcode/day1     3.405s
```

### n cube optimized

```
12:19 $ go test  ./day1/. -bench Fpn -v -benchtime=3s
goos: linux
goarch: amd64
pkg: github.com/Gnoale/adventofcode/day1
BenchmarkFpn
BenchmarkFpn-8          13110172               276 ns/op
PASS
ok      github.com/Gnoale/adventofcode/day1     3.907s
```
### n cube optimized with divide and conquer

```
12:19 $ go test  ./day1/. -bench Fpr -v -benchtime=3s
goos: linux
goarch: amd64
pkg: github.com/Gnoale/adventofcode/day1
BenchmarkFpr
BenchmarkFpr-8          13220934               280 ns/op
PASS
ok      github.com/Gnoale/adventofcode/day1     3.980s
``` 
