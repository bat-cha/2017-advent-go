package day05
import (
	"strings"
	"io/ioutil"
	"strconv"
	"fmt"
)

func Solve(filename string) {

	buf, e := ioutil.ReadFile(filename)

	if e == nil {
		content := string(buf)
		var elements = strings.Split(content,"\n")
		var l int = len(elements)-1
		maze1 := make([]int,l)
		maze2 := make([]int,l)

		for i:=0; i<l; i++ {
			maze1[i],_=strconv.Atoi(elements[i])
			maze2[i]=maze1[i]
		}
		steps1:= Part1(maze1)
		fmt.Println("Part1", steps1," steps to get out of the maze of size ",l)
		steps2:= Part2(maze2)
		fmt.Println("Part2", steps2," steps to get out of the maze of size ",l)

	}

}

func Part1(maze []int) int {
	mazeSize := len(maze)
	position:= 0
	steps :=0
	for position < mazeSize {
		jump := maze[position]
		maze[position]+=1
		position += jump
		steps++
	}
	return steps
}

func Part2(maze []int) int {
	mazeSize := len(maze)
	position:= 0
	steps :=0
	for position < mazeSize {
		jump := maze[position]
		if jump >= 3 {
			maze[position]-=1
		} else {
			maze[position]+=1
		}
		position += jump
		steps++

	}
	return steps
}