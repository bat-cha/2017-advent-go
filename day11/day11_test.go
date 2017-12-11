package day11

import "testing"

func TestPart1(t *testing.T) {
	var s int = 0
	s,_=SolvePuzzle([]string{"ne","ne","ne"})
	if  s!= 3 {
		t.Error("Failure!")
	}
	s,_= SolvePuzzle([]string{"ne","ne","sw","sw"})
	if s!= 0 {
		t.Error("Failure!")
	}
	s,_= SolvePuzzle([]string{"ne","ne","s","s"})
	if s != 2 {
		t.Error("Failure!")
	}
	s,_= SolvePuzzle([]string{"se","sw","se","sw","sw"})
 	if s != 3 {
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