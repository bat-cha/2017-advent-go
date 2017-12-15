package day10
import (
	"strings"
	"io/ioutil"
	"strconv"
	"fmt"
)

func Solve(filename string) {
	SolvePart1(filename,256)
	SolvePart2(filename)
}

func SolvePart1(filename string, listSize int) int {

	var result int = -1

	buf, e := ioutil.ReadFile(filename)

	if e == nil {

		list:= make([]int,listSize)

		for i:=0; i< listSize; i++ {

			list[i]=i
		}

		content := string(buf)
		var elements = strings.Split(strings.TrimSuffix(content,"\n"),",")
		var L int = len(elements)

		lengths := make([]int,L)
		for l:=0; l<L; l++ {
			lengths[l],_=strconv.Atoi(elements[l])

		}

		fmt.Println("input lengths" ,lengths)
		fmt.Println("input list" ,list)
		var currentPosition int = 0
		var skipSize int = 0;

		for l:=0; l< L; l++ {
			//reverse the selected sublist
			for i, j := currentPosition, currentPosition+lengths[l]-1; i < j; i, j = i+1, j-1 {
				list[i%listSize], list[j%listSize] = list[j%listSize], list[i%listSize]

			}
			currentPosition += lengths[l]+skipSize
			skipSize++
		}
		fmt.Println("output list" ,list)
		result = list[0]*list[1]
		fmt.Println("Result : list[0]*list[1]" ,result)
	}
	return result
}

func SolvePart2(filename string) string {

	var result string = ""

	buf, e := ioutil.ReadFile(filename)

	if e == nil {

		content := string(buf)
		result = KnotHash(content)
	}
	return result
}


func DenseHash(sparseHash []int) []int {
	blockSize:=16
	hashSize:= len(sparseHash)/blockSize
	result:=make([]int,hashSize)
	for i:=0;i<hashSize;i++ {
		result[i]=sparseHash[i*blockSize]
		for j:=i*blockSize+1; j<(i+1)*blockSize;j++ {
			result[i] ^= sparseHash[j]
		}
	}
	return result
}

func ToHex(denseHash []int) string {
	var result string = ""
	for i:=0;i<len(denseHash);i++ {
		result+= fmt.Sprintf("%02x",denseHash[i])
	}
	return result
}

func KnotHash(content string) string {

	listSize:=256
	list:= make([]int,listSize)

	for i:=0; i< listSize; i++ {

		list[i]=i
	}

	var L int = len(content)
	suffix:=[5]int{17, 31, 73, 47, 23}

	lengths := make([]int,L+5)
	for l:=0; l<L; l++ {
		lengths[l]=int(content[l])

	}
	for l:=0; l<5;l++ {
		lengths[L+l]=suffix[l]
	}

	//fmt.Println("input lengths" ,lengths)
	//fmt.Println("input list" ,list)
	var currentPosition int = 0
	var skipSize int = 0;
	numRound:=64

	for round:=0; round<numRound;round++ {
		for l := 0; l < L+5; l++ {
			//reverse the selected sublist
			//fmt.Println("start", currentPosition, "end", currentPosition+lengths[l]-1, "end(%)", (currentPosition+lengths[l]-1)%listSize)
			for i, j := currentPosition, currentPosition+lengths[l]-1; i < j; i, j = i+1, j-1 {
				list[i%listSize], list[j%listSize] = list[j%listSize], list[i%listSize]

			}
			currentPosition += lengths[l] + skipSize
			skipSize++
		}
	}
	//fmt.Println("sparse hash", list)
	dense:=DenseHash(list)
	//fmt.Println("dense hash", dense)
	var result = ToHex(dense)
	//fmt.Println("Result in hex", result)
	return result
}
