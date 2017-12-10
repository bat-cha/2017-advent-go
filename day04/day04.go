package day04
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"sort"
)


type empty struct {

}

func Solve(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()
	var validPart1 int = 0
	var validPart2 int = 0

	for e == nil {
		var line string = string(s)
		if (Part1Check(line)) {
			validPart1++
		}
		if (Part2Check(line)) {
			validPart2++
		}
		s, _, e = r.ReadLine()
	}
	fmt.Println("part1 valid ",validPart1)
	fmt.Println("part2 valid ",validPart2)

}

func Part1Check(line string) bool {
	var elements = strings.Split(line, " ")
	words := map[string]empty{}
	var l int = len(elements)

	for i := 0; i < l; i++ {
		word := elements[i]
		_, present := words[word]
		if present {
			return false
		} else {
			words[word] = empty{}
		}
	}
	return true
}

func Part2Check(line string) bool {
	var elements = strings.Split(line, " ")
	words := map[string]empty{}
	var l int = len(elements)

	for i := 0; i < l; i++ {
		word := SortString(elements[i])
		_, present := words[word]
		if present {
			return false
		} else {
			words[word] = empty{}
		}
	}
	return true
}


type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}