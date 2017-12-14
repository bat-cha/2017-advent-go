package day12
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type empty struct {}

func Solve(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()

	var totalGroupCount int = 0
	var groupId = "0"
	groups :=map[string]map[string]empty{groupId:map[string]empty{groupId:empty{}}}
	buffer:=map[string][]string{}
	var unlinked int =42

	for unlinked >1 {

		for e == nil {
			var line string = string(s)
			var elements= strings.Split(line, " <-> ")
			pid := elements[0]
			_, inGroup := groups[groupId][pid]
			links := strings.Split(elements[1], ", ")
			if inGroup {
				groups[groupId][pid] = empty{}
				for _, p := range buffer[pid] {
					groups[groupId][p] = empty{}
				}
				delete(buffer, pid)
				for _, p := range links {
					groups[groupId][p] = empty{}
					for _, lp := range buffer[p] {
						groups[groupId][lp] = empty{}
					}
					delete(buffer, p)
				}
			} else {
				buffer[pid] = links
			}

			s, _, e = r.ReadLine()
		}

		var linkFound = true
		for linkFound {
			linkFound = false
			for p, links := range buffer {
				_, found := groups[groupId][p]
				if found {
					for _, l := range links {
						groups[groupId][l] = empty{}
					}
					delete(buffer, p)
					linkFound = linkFound || true
				}
			}

		}

		fmt.Println(len(groups[groupId]), len(buffer))
		var key string
		for key, _ = range buffer {
			break
		}
		groupId=key
		groups[groupId]=map[string]empty{groupId:empty{}}
		unlinked=len(buffer)
		totalGroupCount++
	}
	totalGroupCount++
	fmt.Println(totalGroupCount)


}