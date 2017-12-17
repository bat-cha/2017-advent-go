package day17
import (
	"io/ioutil"
	"log"
	"strconv"
	"fmt"
)


func Solve(filename string) {

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	} else {
		steps,_:=strconv.Atoi(string(buf))
		SolvePart1(steps)
		//SolvePart2Naive(steps)
		SolvePart2(steps,50000000)
	}
}

func SolvePart1(steps int) int {
	return SolvePuzzle(steps,2017,2017)

}


func SolvePart2Naive(steps int) int {
	return SolvePuzzle(steps,50000000,0)

}


func SolvePuzzle(steps int, iteration int, valueAfter int) int {

	var result int = 0

	fmt.Println("#steps:",steps)

	buffer :=[]int{0}
	var currentPosition int = 0
	var stepForwardCount int = 0

	for stepForwardCount < iteration+1 {
		stepForwardCount++
		currentPosition = (currentPosition +  steps) % stepForwardCount

		if (currentPosition == 0) {
			fmt.Println("insert after 0 for",stepForwardCount)
		}
		buffer = append(buffer, 0)
		currentPosition++
		copy(buffer[currentPosition+1:], buffer[currentPosition:])
		buffer[currentPosition] = stepForwardCount

		if (stepForwardCount % 10000 ) == 0 {
			fmt.Println(stepForwardCount)
		}
	}

	search: for i,v:=range buffer {
		if v == valueAfter {
			result = buffer[(i+1)%stepForwardCount]
			break search
		}
	}
	fmt.Println(result)
	return result
}



func SolvePart2(steps int, iteration int,) int {

	var result int = 0

	fmt.Println("#steps:",steps)

	var currentPosition int = 0
	var stepForwardCount int = 0

	for stepForwardCount < iteration+1 {
		stepForwardCount++
		currentPosition = (currentPosition +  steps) % stepForwardCount

		if currentPosition == 0 {
			fmt.Println("insert after 0 for",stepForwardCount)
			result=stepForwardCount
		}
		currentPosition++

	}

	fmt.Println(result)
	return result
}


func PrintBuffer(bufffer []int) string {
	var result string
	for i:=0;i<len(bufffer);i++ {
		result+= fmt.Sprint(" ",bufffer[i])
	}
	return result
}