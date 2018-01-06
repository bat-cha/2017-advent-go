package day23
import (
	"testing"
)


func TestPuzzle(t *testing.T) {
	Solve("input.txt")
}


func BenchmarkPuzzle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Solve("input.txt")
	}
}
