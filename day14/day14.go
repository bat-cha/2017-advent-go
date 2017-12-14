package day13
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func Solve(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()

	scanners:=map[int]int{}

	for e == nil {
		var line string = string(s)
		var elements= strings.Split(line, ": ")
		depth,_ := strconv.Atoi(elements[0])
		layerRange,_ :=  strconv.Atoi(elements[1])
		scanners[depth]=layerRange
		s, _, e = r.ReadLine()
	}

	//var part1Severity = 0;
	var part2Delay= 0;
	var caught bool = true;

	for caught {
		caught=false;
		for depth,layerRange:=range scanners {
			if ((depth+part2Delay)%(2*(layerRange-1))) == 0 {
				//part1Severity+=depth*layerRange
				caught=true
				part2Delay++
				break
			}
		}
	}

	fmt.Println("part2 Minimal delay :",part2Delay)

}