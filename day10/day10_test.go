package day10
import (
	"testing"
)

func TestPart1(t *testing.T) {
	if SolvePart1("test_input.txt",5) != 12 {
		t.Error("Failure")
	}
}

func TestPart2(t *testing.T) {

	if SolvePart2("test_input2.txt",256) != "a2582a3a0e66e6e86e3812dcb672a272" {
		t.Error("Failure")
	}
	if SolvePart2("test_input3.txt",256) != "33efeb34ea91902bb2f59c9920caa6cd" {
		t.Error("Failure")
	}
	if SolvePart2("test_input4.txt",256) != "3efbe78a8d82f29979031a4aa0b16a9d" {
		t.Error("Failure")
	}
	if SolvePart2("test_input5.txt",256) != "63960835bcdc130f0b66d7ff4f6a5a8e" {
		t.Error("Failure")
	}

}

func TestDenseHash(t *testing.T) {
	test:=[]int{65,27,9,1,4,3,40,50,91,7,6,0,2,5,68,22}
	if DenseHash(test)[0]!=64 {
		t.Error("Failure")
	}
}

func TestToHex(t *testing.T) {
	test:=[]int{64, 7, 255}
	if ToHex(test)!="4007ff" {
		t.Error("Failure",ToHex(test))
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

