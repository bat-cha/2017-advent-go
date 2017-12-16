package day16
import (
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"strconv"
	"unicode/utf8"
)

type DanceStage struct {
	size int
	indicesMove []int
	partnerSwap map[rune]rune

}

type DanceMove interface {
	dance(stage *DanceStage)
}

type Spin struct {
	size int
}

type Exchange struct {
	position1 int
	position2 int
}

type Partner struct {
	program1 rune
	program2 rune
}

func (s Spin) dance(stage *DanceStage) {
	newIm:=make([]int,stage.size)

	for i:=stage.size-s.size;i<stage.size;i++ {
		newIm[i-stage.size+s.size]= stage.indicesMove[i]
	}
	for i:=0;i<stage.size-s.size;i++ {
		newIm[s.size+i]=stage.indicesMove[i]
	}
	for i,v:=range newIm {
		stage.indicesMove[i]=v
	}
}

func (e Exchange) dance(stage *DanceStage) {
	stage.indicesMove[e.position1], stage.indicesMove[e.position2] = stage.indicesMove[e.position2],stage.indicesMove[e.position1]
}

func (p Partner) dance(stage *DanceStage) {
	newP:=map[rune]rune{}
	for key,value:=range stage.partnerSwap {
		if value == p.program1 {
			newP[key]=p.program2
		}
		if value == p.program2 {
			newP[key]=p.program1
		}
	}
	for k,v:=range newP {
		stage.partnerSwap[k]=v
	}
}

func Solve(filename string) {

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	} else {
		moves:=strings.Split(string(buf),",")
		SolvePart1(moves)
		SolvePart2(moves,1000000000)
	}
}

func SolvePart1(moves []string) string {

	return SolveDance(moves, 16, 1)
}

func SolvePart2(moves []string,iterations int64) string {

	return SolveDance(moves, 16, iterations)
}

func SolveDance(moves []string, size int, iterations int64) string {

	programs:=[]rune("abcdefghijklmnop")
	indicesMove:=make([]int,size)
	partnerSwap:=map[rune]rune{}

	for i:=0;i<size;i++ {
		indicesMove[i]=i
		partnerSwap[programs[i]]=programs[i]
	}
	stage:=DanceStage{size,indicesMove,partnerSwap}
	for _, move := range moves {
		m:=buildDanceMove(move)
		m.dance(&stage)

	}
	var iteration int64 = iterations
	for iteration > 0 {
		if (iteration & 1) == 1 {
			newProgram:=make([]rune,size)
			for i,v:=range stage.indicesMove {
				newProgram[i] = stage.partnerSwap[programs[v]]
			}
			for i,v:=range newProgram {
				programs[i]=v
			}
		}

		newIm:=make([]int,size)
		for i,m:=range stage.indicesMove {
			newIm[i] = stage.indicesMove[m]
		}
		stage.indicesMove = newIm
		newPS:=map[rune]rune{}
		for k,v:= range stage.partnerSwap {
			newPS[k]=stage.partnerSwap[v]
		}
		stage.partnerSwap = newPS

		iteration >>= 1
	}

	var result string =string(programs[0:size])
	fmt.Println(result)
	return result
}


func buildDanceMove(move string) DanceMove {

	var result DanceMove

	switch string(move[0]) {
	case "s":
		spin,e:=strconv.Atoi(move[1:])
		if e!=nil {
			log.Fatal(e)
		}
		result = Spin{spin}

	case "x":
		positions:=strings.Split(move[1:],"/")
		position1,_:=strconv.Atoi(positions[0])
		position2,_:=strconv.Atoi(positions[1])
		result = Exchange{position1,position2}

	case "p":
		programs:=strings.Split(move[1:],"/")
		program1,_:=utf8.DecodeRuneInString(programs[0])
		program2,_:=utf8.DecodeRuneInString(programs[1])
		result = Partner{program1,program2}

	default:
		result = nil
	}
	return result
}

