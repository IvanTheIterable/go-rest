package main

import (
	"fmt"
	"testing"
)

func BenchmarkCodegen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("abc")
	}
}
