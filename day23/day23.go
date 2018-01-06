package day23
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

func Solve(filename string) {
	SolvePart1(filename)
	//SolvePart2Naive(filename)
	SolvePart2()
}

func SolvePart1(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()


	var instructions []string
	for e == nil {
		instructions=append(instructions,string(s))
		s, _, e = r.ReadLine()
	}

	registers:=map[string]int64{}
	var currentInstruction int64 = 0
	var mul bool
	var mulCnt int64 =0
	for currentInstruction>=0 && currentInstruction<int64(len(instructions))  {
		instruction:=instructions[currentInstruction]
		currentInstruction,mul = RunInstruction(instruction,registers,currentInstruction)
		if mul {
			mulCnt++
		}
	}
	fmt.Println("mul run ",mulCnt)
}

func SolvePart2Naive(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()


	var instructions []string
	for e == nil {
		instructions=append(instructions,string(s))
		s, _, e = r.ReadLine()
	}

	registers:=map[string]int64{"a":1}
	var currentInstruction int64 = 0
	var count int =0
	for currentInstruction>=0 && currentInstruction<int64(len(instructions)) && count<100  {
		instruction:=instructions[currentInstruction]
		currentInstruction,_ = RunInstruction(instruction,registers,currentInstruction)
		count++
	}
	fmt.Println("h register ",registers["h"])

}


func SolvePart2() {
	h:=0
	c:=126300
	for b:=109300;b<c+1;b=b+17 {
		var incrementH bool = false
		for d:=2;d<int(math.Sqrt(float64(b)));d++ {
			if (b%d)==0 {
				incrementH=true
				break
			}
		}
		if incrementH {
			h++
		}
		incrementH=false
	}
	fmt.Println("h register ",h)

}

func RunInstruction(instruction string, registers map[string]int64, currentInstruction int64) (int64,bool) {
	arguments:=strings.Split(instruction," ")
	cmd:= arguments[0]
	var result int64 = currentInstruction + 1
	var mul bool = false
	switch cmd {
	case "set":
		x:=arguments[1]
		y:=arguments[2]
		//fmt.Println(cmd,x,y)
		val:=GetValue(y,registers)
		registers[x]=val
	case "sub":
		x:=arguments[1]
		y:=arguments[2]
		//fmt.Println(cmd,x,y)
		val1:=GetValue(x,registers)
		//fmt.Println(val1)
		val2:=GetValue(y,registers)
		registers[x]=val1-val2
	case "mul":
		x:=arguments[1]
		y:=arguments[2]
		//fmt.Println(cmd,x,y)
		val1:=GetValue(x,registers)
		val2:=GetValue(y,registers)
		registers[x]=val1*val2
		mul=true
	case "jnz":
		x:=arguments[1]
		y:=arguments[2]
		//fmt.Println(cmd,x,y)
		val1:=GetValue(x,registers)
		if val1 != 0 {
			val2:=GetValue(y,registers)
			result=currentInstruction+val2
		}
	}
	return result,mul
}

func GetValue(register string, registers map[string]int64) int64 {

	parsedValue,e:=strconv.Atoi(register)
	var value int64 = int64(parsedValue)
	if e!= nil {
		v,present:=registers[register]
		if !present {
			value=0
			registers[register]=int64(value)
		} else {
			value=v
		}
	}
	return value
}