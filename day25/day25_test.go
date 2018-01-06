package day25
import (
	"testing"
)

func TestInput(t *testing.T) {
	if Solve("test_input.txt") != 3 {
		t.Error("Failed!")
	}
}

func TestPuzzle(t *testing.T) {
	Solve("input.txt")
}

func BenchmarkPuzzle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Solve("input.txt")
	}
}
