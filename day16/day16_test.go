package day16
import (
	"testing"
)

//
func TestInput(t *testing.T) {
	moves:= []string{"s1","x3/4","pe/b","pb/e","pe/b"}
	SolveDance(moves,5,1)
}

func TestInput2(t *testing.T) {
	moves := []string{"s1", "x3/4", "pe/b"}
	SolveDance(moves, 5, 2)
}

func TestPuzzle(t *testing.T) {
	Solve("input.txt")
}

func BenchmarkPuzzle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Solve("input.txt")
	}
}
