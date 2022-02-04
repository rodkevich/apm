package main

import (
	"strings"
	"testing"
)

func DNAtoRNA(dna string) string {
	return strings.Replace(dna, "T", "U", -1)
}

func DNAtoRNA2(dna string) string {

	const shouldBeChanged = "T"
	rtn := ""

	if dna == rtn {
		return dna
	}

	for _, char := range dna {
		if string(char) == shouldBeChanged {
			rtn += "U"
			continue
		}
		rtn += string(char)
	}
	return rtn
}

func BenchmarkDNAtoRNA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DNAtoRNA("GCAT")
	}
}

func BenchmarkDNAtoRNA2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DNAtoRNA2("GCAT")
	}
}
