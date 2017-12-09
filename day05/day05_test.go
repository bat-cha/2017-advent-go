package day05
import (
	"testing"
)

func TestPart1(t *testing.T) {
	test1 := []int{0,3,0,1,-3}
	if Part1(test1)!=5 {
		t.Error("Failure!")
	}
}

func TestPart2(t *testing.T) {
	test2 := []int{0,3,0,1,-3}
	if Part2(test2)!=10 {
		t.Error("Failure!")
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