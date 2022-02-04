package main

import (
	"strings"
	"testing"
	"unicode"
)

func BenchmarkMyString_IsUpperCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsUpperCase("asdsdadsqweqeQEWdasdADASDqweqwedasdCXACasdsdadsqweqeQEWdasdADASDqweqwedasdCXACasdsdadsqweqeQEWdasdADASDqweqwedasdCXAC")
	}
}

func IsUpperCase(s string) bool {
	return strings.ToUpper(s) == s

}

func BenchmarkMyString_IsUpperCase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsUpperCase2("asdsdadsqweqeQEWdasdADASDqweqwedasdCXACasdsdadsqweqeQEWdasdADASDqweqwedasdCXACasdsdadsqweqeQEWdasdADASDqweqwedasdCXAC")
	}
}

func IsUpperCase2(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return false
		}
	}
	return true
}
