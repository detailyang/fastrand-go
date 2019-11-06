<p align="center">
  <b>
    <span style="font-size:larger;">fastrand-go</span>
  </b>
  <br />
   <a href="https://travis-ci.org/detailyang/fastrand-go"><img src="https://travis-ci.org/detailyang/fastrand-go.svg?branch=master" /></a>
   <a href="https://ci.appveyor.com/project/detailyang/fastrand-go"><img src="https://ci.appveyor.com/api/projects/status/hbpj944ankoy9sh5?svg=true" /></a>
   <a href="https://godoc.org/github.com/detailyang/fastrand-go">
      <img src="https://godoc.org/github.com/detailyang/fastrand-go?status.svg"/>
   </a>
   <br />
   <b>Package fastrand exports runtime.fastrand to public</b>
</p>

````bash
╰─λ go test -benchmem -run="^\$" -bench Benchmark ./... -v                                                                          306ms
goos: darwin
goarch: amd64
pkg: github.com/detailyang/fastrand-go
BenchmarkFastRand-8   	493277304	         2.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkMathRand-8   	54207786	        19.2 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/detailyang/fastrand-go	2.502s
````