package day22
import (
	"testing"
)


func TestInputPart1(t *testing.T) {
	virus,infectedNodes:=initialize("test_input.txt")
	if SolvePart1(virus,7,infectedNodes) !=5 {
		t.Error("Failed!")
	}
	virus,infectedNodes=initialize("test_input.txt")
	if SolvePart1(virus,70,infectedNodes) !=41 {
		t.Error("Failed!")
	}
	virus,infectedNodes=initialize("test_input.txt")
	if SolvePart1(virus,10000,infectedNodes) !=5587 {
		t.Error("Failed!")
	}
}

func TestInputPart2(t *testing.T) {
	virus,infectedNodes:=initialize("test_input.txt")
	if SolvePart2(virus,100,infectedNodes) !=26 {
		t.Error("Failed!")
	}
	virus,infectedNodes=initialize("test_input.txt")
	if SolvePart2(virus,10000000,infectedNodes) !=2511944 {
		t.Error("Failed!")
	}
}

func TestPuzzle(t *testing.T) {
	Solve("input.txt")
}
//
//func BenchmarkPuzzle(b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		Solve("input.txt")
//	}
//}
