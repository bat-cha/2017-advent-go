package day14
import (
	"testing"
	"../day10"
)

func TestHexToBin(t *testing.T) {

	if day10.KnotHash("AoC 2017")!="33efeb34ea91902bb2f59c9920caa6cd" {
		t.Error("Failed")
	}

	if HexToBin("0") != "0000" {
		t.Error("Failed")
	}
	if HexToBin("1") != "0001" {
		t.Error("Failed")
	}
	if HexToBin("e") != "1110" {
		t.Error("Failed")
	}
	if HexToBin("f") != "1111" {
		t.Error("Failed")
	}
	if HexToBin("ef") != "11101111" {
		t.Error("Failed")
	}
	if HexToBin("a0c2017") != "1010000011000010000000010111" {
		t.Error("Failed")
	}
}

func TestInput(t *testing.T) {
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
