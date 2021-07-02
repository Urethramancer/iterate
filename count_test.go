//go:generate go test -v -bench=. ./... -count=1 -benchmem
package iterate_test

import (
	"fmt"
	"testing"

	iterate "github.com/Urethramancer/count"
)

func TestCount(t *testing.T) {
	for _, x := range iterate.Count(10) {
		if x > 0 {
			fmt.Printf(", ")
		}
		fmt.Printf("%d", x)
	}

	println("\n")
	for _, x := range iterate.Count(1, 10) {
		if x > 1 {
			fmt.Printf(", ")
		}
		fmt.Printf("%d", x)
	}

	println("\n")
	for _, x := range iterate.Count(2, 10, 2) {
		if x > 2 {
			fmt.Printf(", ")
		}
		fmt.Printf("%d", x)
	}
	println("\n")

	for _, x := range iterate.Count(10, 1, -1) {
		if x < 10 {
			fmt.Printf(", ")
		}
		fmt.Printf("%d", x)
	}
	println("")
}

func BenchmarkCount1(b *testing.B) {
	y := 0
	for _, x := range iterate.Count(9) {
		y += x
	}
}

func BenchmarkCount2(b *testing.B) {
	y := 0
	for _, x := range iterate.Count(1, 10) {
		y += x
	}
}

func BenchmarkCount3(b *testing.B) {
	y := 0
	for _, x := range iterate.Count(2, 10, 2) {
		y += x
	}
}

func BenchmarkNegative(b *testing.B) {
	y := 0
	for _, x := range iterate.Count(10, 1, -1) {
		y += x
	}
}
