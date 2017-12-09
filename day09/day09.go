package day09
import (
	"strings"
	"fmt"
	"bufio"
	"io"
	"log"
	"os"
)

func Solve(filename string) (int,int) {
	f, e:= os.Open(filename)
	if e == nil {
		return SolvePuzzle(f)
	}
	return -1,-1
}

func SolveString(content string) (int, int) {
	fmt.Println(content)
	return SolvePuzzle(strings.NewReader(content))
}

func SolvePuzzle(reader io.Reader) (int, int) {
	var score int = 0
	var nonCanceled int =0
	groupsDepth:=map[int]int{}
	var currentDepth int =0
	var ignore bool = false
	var ignoreOnce bool = false
	r := bufio.NewReader(reader)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if !ignoreOnce {
				current:=string(c)
				//fmt.Println(current)
				if ignore && (current != "!"){
					nonCanceled++
				}
				if ignore && (current == ">") {
					//fmt.Println("end of ignore")
					nonCanceled--
					ignore = false
				} else {
					if ignore  {
						if current == "!" {
							ignoreOnce=true
							//fmt.Println("start of ignoreOnce")
						}
					} else {
						switch current {
							case "<":
								ignore=true
								//fmt.Println("start of ignore")
							case "!":
								ignoreOnce=true
								//fmt.Println("start of ignoreOnce")
							case "{":
								currentDepth++
							case "}":
								d:=getDepth(currentDepth,groupsDepth)
								groupsDepth[currentDepth]=d+1
								currentDepth--
							default:



						}
					}
				}
			} else {
				//fmt.Println("ignored ", string(c))
				ignoreOnce = false
				//fmt.Println("end of ignoreOnce")
			}

		}
	}
	for depth,count := range groupsDepth {
		score+=depth*count
	}
	fmt.Println("Score:",score, "Non Canceled:",nonCanceled)
	return score, nonCanceled
}


func getDepth(currentDepth int, groupsDepth map[int]int) int {
	var depth int
	var present bool
	depth,present=groupsDepth[currentDepth]
	if !present {
		depth=0
	}
	return depth
}