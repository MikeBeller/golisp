# golisp

Building a simple LISP in go, just for fun.

Interesting thing learned: Recursive go interfaces have a
recursive notion of equality.   This means that in my
implementation of LISP, list1 == list2 checks the full 
list for equality, not just the first element!

See equal_bench_test.go

mike@hal:~/github/golisp$ go test -bench=List
goos: linux
goarch: amd64
BenchmarkList100-4           	 1000000	      1235 ns/op
BenchmarkList1k-4            	  100000	     13268 ns/op
BenchmarkList10k-4           	   10000	    141987 ns/op
BenchmarkList10kDiff-4       	   10000	    143047 ns/op
BenchmarkList10kDiffSoon-4   	10000000	       128 ns/op
PASS

You can see that for a list of 10k it takes 10 times as long to
do == as a list of 1k, which takes 10 times as long as a list of 100.

Also notice that if the two lists differ early on, the comparison
is fast (List10kDiffSoon)

