# skido
Sample Golang project to showcase Golang's producer consumer pattern and Goroutine's worker pool


1.Steps to run your program:-
go run main.go

2.Steps to run test/benchmarks of your program:-
go test -run=. -bench=. -benchtime=5s -count 5 -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out . | tee bench.txt

3.Compare the results of benchmark:-
goos: linux
goarch: amd64
pkg: github.com/skido
BenchmarkPartA-3   	       1	19023619488 ns/op	      64 B/op	       1 allocs/op
BenchmarkPartA-3   	       1	19022072512 ns/op	      64 B/op	       1 allocs/op
BenchmarkPartA-3   	       1	19017470522 ns/op	      64 B/op	       1 allocs/op
BenchmarkPartA-3   	       1	19033760569 ns/op	      64 B/op	       1 allocs/op
BenchmarkPartA-3   	       1	19029968588 ns/op	      64 B/op	       1 allocs/op
BenchmarkPartB-3   	       1	11239389900 ns/op	    5664 B/op	      61 allocs/op
BenchmarkPartB-3   	       1	11239711460 ns/op	    6816 B/op	      65 allocs/op
BenchmarkPartB-3   	       1	11240144347 ns/op	    7200 B/op	      65 allocs/op
BenchmarkPartB-3   	       1	11244023013 ns/op	    1824 B/op	      51 allocs/op
BenchmarkPartB-3   	       1	11247076131 ns/op	    1824 B/op	      51 allocs/op
PASS
ok  	github.com/skido	302.649s

4.Compare the results of profiling:-
Refer to images uploaded in the repo for comparison (profile001.png and profile002.png)
