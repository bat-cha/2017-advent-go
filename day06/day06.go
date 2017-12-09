package day06
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
		var elements = strings.Split(strings.TrimSuffix(content,"\n"),"\t")

		var l int = len(elements)
		fmt.Println(elements,l)
		blocks1 := make([]int,l)
		for i:=0; i<l; i++ {
			blocks1[i],_=strconv.Atoi(elements[i])

		}
		cycles1 := Part1(blocks1,l)
		fmt.Println("Part1", cycles1," cycles to see a known distribution for blocks of size",l)


	}

}

type count struct {
	cycle int

}

func Part1(blocks []int, size int) int {
	fmt.Println(blocks)
	cycles:=0
	newState:= true
	states := map[string]count {}
	var state string
	for newState {
		state = Distribute(blocks,size, findMaxBlockIndex(blocks,size))
		_,present := states[state]
		cycles++
		newState = !present
		if !present {
			states[state] = count{cycles}
		}
		//fmt.Println(state)
	}
	fmt.Println("duplicate state " , state)
	fmt.Println("cycle length " , cycles-states[state].cycle)
	return cycles
}

func findMaxBlockIndex(blocks []int, size int) int {
	max:=0
	maxIndex :=0
	for i:= 0; i<size; i++ {
		if blocks[i]>max {
			max = blocks[i]
			maxIndex = i
		}
	}
	return maxIndex

}

func Distribute(blocks []int, size int, index int) string {

	value:= blocks[index]
	blocks[index]=0
	for i:=index+1; value>0; i++ {
		blocks[i%size]+=1
		value--
	}
	return BlockToKey(blocks,size)
}

func BlockToKey(blocks []int, size int) string {
	key := make([]string,size)
	for i:=0;i<size;i++ {
		key[i]=strconv.Itoa(blocks[i])
	}
	return strings.Join(key,"-")
}
