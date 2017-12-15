package day15
import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
	"log"
)
const dividor int64 = 2147483647
const factorA int64 = 16807
const factorB int64 = 48271

type empty struct {

}

func Solve(filename string) {

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	} else {
		lines:=strings.Split(string(buf),"\n")
		genAstart,_:=strconv.ParseInt(strings.Split(lines[0],"Generator A starts with ")[1],10,64)
		genBstart,_:=strconv.ParseInt(strings.Split(lines[1],"Generator B starts with ")[1],10,64)
		SolvePart1(genAstart,genBstart)
		SolvePart2(genAstart,genBstart)
	}
}

func SolvePart1(genAStart int64, genBStart int64) int64 {

	var result int64 = 0
	var pairs int = 0
	var previousA int64 = genAStart
	var previousB int64 = genBStart

	for pairs < 40000000 {

		currentA:=(previousA*factorA)%dividor
		currentB:=(previousB*factorB)%dividor
		previousA=currentA
		previousB=currentB
		pairs++
		//fmt.Println(currentA,currentB)


		binA:=fmt.Sprintf("%032b",currentA )
		binB:=fmt.Sprintf("%032b",currentB )

		runesA := []rune(binA)
		runesB := []rune(binB)

		if string(runesA[16:32]) == string(runesB[16:32]) {
			result++
		}

		//fmt.Println(binA)
		//fmt.Println(binB)


	}

	fmt.Println(genAStart,genBStart)
	fmt.Println("Part 1 Result ",result)

	return result
}


func SolvePart2(genAStart int64, genBStart int64) int64 {

	var result int64 = 0
	var pairs int = 0
	var previousA int64 = genAStart
	var previousB int64 = genBStart

	for pairs < 5000000 {

		var validA bool = false
		var currentA int64 = 0
		for !validA {
			currentA =(previousA*factorA)%dividor
			previousA=currentA
			validA = (currentA % 4) == 0
		}
		var validB bool = false
		var currentB int64 =0
		for !validB {
			currentB=(previousB*factorB)%dividor
			previousB=currentB
			validB = (currentB % 8) == 0
		}

		pairs++
		//fmt.Println(currentA,currentB)


		binA:=fmt.Sprintf("%032b",currentA )
		binB:=fmt.Sprintf("%032b",currentB )

		runesA := []rune(binA)
		runesB := []rune(binB)

		if string(runesA[16:32]) == string(runesB[16:32]) {
			result++
		}

		//fmt.Println(binA)
		//fmt.Println(binB)


	}

	fmt.Println(genAStart,genBStart)
	fmt.Println("Part 2 Result ",result)

	return result
}
