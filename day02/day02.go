package day02
import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"strings"
)

func Solve(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()
	var checksum int = 0
	for e == nil {
		var line string = string(s)
		var elements = strings.Split(line,"\t")
		var l int = len(elements)
		var n int = 0
		var m int = 0
		//n, err = strconv.Atoi(elements[0])
		//m, err = strconv.Atoi(elements[1])
		//var max int = n
		//var min int = n

		Search: for i:= 0; i < l; i++ {
			for j:= i+1; j<l; j++ {
				n, err = strconv.Atoi(elements[i])
				m, err = strconv.Atoi(elements[j])
				fmt.Println(n,m)
				if n%m == 0 {
					fmt.Println("found", n,m)
					checksum += n/m
					break Search
				} else if m%n ==0 {
					fmt.Println("found", m,n)
					checksum += m/n
					break Search
				}
			}
		//	min = Min(n,min)
		//	max = Max(n,max)
		}

		//fmt.Println(min,max)
		//checksum += (max - min)
		fmt.Println(checksum)
		s, _, e = r.ReadLine()
	}
	fmt.Println(checksum)

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


func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}