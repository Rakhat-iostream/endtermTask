package main

import (
	"io/ioutil"
	"testing"
)

// go test -bench . -benchmem
func BenchmarkWordEnumerator(b *testing.B) {
	WordEnumerator(ioutil.Discard)
}
