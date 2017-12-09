package day06
import (
	"testing"
)

func TestPart1(t *testing.T) {
	test1 := []int{0,2,7,0}
	if Part1(test1,4)!=5 {
		t.Error("Failure!")
	}
	test2 := []int{1,3,0,0}
	if Part1(test2,4)!=10 {
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

