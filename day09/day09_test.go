package day09
import (
	"testing"
)

func TestPart1(t *testing.T) {
	garbage:= [7]string{"<>","<random characters>","<<<<>","<{!>}>","<!!>","<!!!>>","<{o\"i!a,<{i<a>"}
	for _,content:= range garbage {
		v,_ := SolveString(content)
		if v != 0 {
			t.Error("Failed for input ",content)
		}
	}
	groups := map[string]int{"{}":1, "{{{}}}":6, "{{},{}}":5, "{{{},{},{{}}}}":16, "{<a>,<a>,<a>,<a>}":1, "{{<ab>},{<ab>},{<ab>},{<ab>}}":9,"{{<!!>},{<!!>},{<!!>},{<!!>}}":9,"{{<a!>},{<a!>},{<a!>},{<ab>}}":3}

	for content,score := range groups {
		calculatedScore,_:=SolveString(content)
		if calculatedScore != score {
			t.Error("Failed for input ",content, " expected ",score, " got ",calculatedScore)
		} else {
			t.Log("Success for input ",content, " expected ",score, " got ",calculatedScore)
		}
	}

}

func TestPart2(t *testing.T) {
	garbage := map[string]int{"<>":0, "<random characters>":17, "<<<<>":3, "<{!>}>":2, "<!!>":0, "<!!!>>":0, "<{o\"i!a,<{i<a>":10}
	for content,nonCanceled := range garbage {
		_,calculatedNonCanceled:=SolveString(content)
		if calculatedNonCanceled != nonCanceled {
		t.Error("Failed for input ",content, " expected ",nonCanceled, " got ",calculatedNonCanceled)
	} else {
		t.Log("Success for input ",content, " expected ",nonCanceled, " got ",calculatedNonCanceled)
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

