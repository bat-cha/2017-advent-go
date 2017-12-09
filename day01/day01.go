package day01
import (
	"fmt"
  	"log"
	"io/ioutil"
	"strconv"
)

func Solve(filename string) {

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	} else {
		var sum int
		var sequence []int = ConvertToI(buf)
		sum = FindSumPart1(sequence)
		fmt.Println("Part 1", sum)
		sum = FindSumPart2(sequence)
		fmt.Println("Part 2", sum)
	}
}

func ConvertToI(buf []byte) ([]int) {

  var l int = len(buf)-1
  var res []int = make([]int,l)
  for i:= 0; i < l; i++ {
  	var v int
  	v, err := strconv.Atoi(string(buf[i]))
  	if err == nil {
  		res[i] = v
	}
  }
  return res
}

func FindSumPart1(sequence []int) (int) {
	return FindSumPart(sequence,1)
}
func FindSumPart2(sequence []int) (int) {
	return FindSumPart(sequence,len(sequence)/2)
}

func FindSumPart(sequence []int, step int) (int) {
	var l int = len(sequence)
	var sum int = 0
	for i := 0; i < l; i++ {
		if sequence[i] == sequence[(i+step)%l] {
			sum += sequence[i]
		}
	}
	return sum
}
