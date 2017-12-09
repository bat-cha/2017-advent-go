package day04
import (
	"testing"
)

func TestPart1(t *testing.T) {

	cases := map[string]bool{"aa bb cc dd ee":true, "aa bb cc dd aa":false, "aa bb cc dd aaa":true}

	for input,valid := range cases {
		if Part1Check(input) != valid {
			t.Error("Failure for",input)
		}
	}

}

func TestPart2(t *testing.T) {

	cases := map[string]bool{"abcde fghij":true, "abcde xyz ecdab":false, "a ab abc abd abf abj":true, "iiii oiii ooii oooi oooo":true, "oiii ioii iioi iiio":false}

	for input,valid := range cases {
		if Part2Check(input) != valid {
			t.Error("Failure for",input)
		}
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