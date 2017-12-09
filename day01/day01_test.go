package day01
import (
	"testing"
)

func TestFindSumPart1(t *testing.T) {
	var sum1 int = FindSumPart1([]int{1,1,2,2})
	if sum1 != 3 {
		t.Error("expected 3, got ",sum1)
	}
	var sum2 int = FindSumPart1([]int{1,1,1,1})
	if sum2 != 4 {
		t.Error("expected 4, got ",sum2)
	}
	var sum3 int = FindSumPart1([]int{1,2,3,4})
	if sum3 != 0 {
		t.Error("expected 0, got ",sum3)
	}
	var sum4 int = FindSumPart1([]int{9,1,2,1,2,1,2,9})
	if sum4 != 9 {
		t.Error("expected 9, got ",sum4)
	}

}


func TestFindSumPart2(t *testing.T) {
	var sum1 int = FindSumPart2([]int{1,2,1,2})
	if sum1 != 6 {
		t.Error("expected 6, got ",sum1)
	}
	var sum2 int = FindSumPart2([]int{1,2,2,1})
	if sum2 != 0 {
		t.Error("expected 0, got ",sum2)
	}
	var sum3 int = FindSumPart2([]int{1,2,3,4,2,5})
	if sum3 != 4 {
		t.Error("expected 4, got ",sum3)
	}
	var sum4 int = FindSumPart2([]int{1,2,3,1,2,3})
	if sum4 != 12 {
		t.Error("expected 12, got ",sum4)
	}
	var sum5 int = FindSumPart2([]int{1,2,1,3,1,4,1,5})
	if sum5 != 4 {
		t.Error("expected 4 got ",sum5)
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