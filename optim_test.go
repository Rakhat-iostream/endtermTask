package main

import (
	"io/ioutil"
	"testing"
)

// go test -bench . -benchmem
func BenchmarkWordCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReadFile(ioutil.Discard)
	}
}
