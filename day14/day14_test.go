package day13
import (
	"testing"
)

func TestPart1(t *testing.T) {
	Solve("test_input.txt")
}

func TestPuzzle(t *testing.T) {
	Solve("input.txt")
}

func BenchmarkPuzzle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Solve("input.txt")
	}
}
