package day18
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"time"
)

const lastSoundPlayed string = "last"
const sndCounter string = "sndCount"
const deadlockTimeout time.Duration = 2

func Solve(filename string) {
	SolvePart1(filename)
	SolvePart2(filename)
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

	registers:=map[string]int{}
	var currentInstruction int = 0
	var firstRcv bool = false
	for currentInstruction>=0 && currentInstruction<len(instructions) && !firstRcv {
		instruction:=instructions[currentInstruction]
		currentInstruction,firstRcv = RunInstruction(instruction,registers,currentInstruction)
	}
}


func SolvePart2(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()


	instructions:=[]string{}
	for e == nil {
		instructions=append(instructions,string(s))
		s, _, e = r.ReadLine()
	}

	registers0:=map[string]int{"p":0}
	registers1:=map[string]int{"p":1}
	var currentInstruction0 int = 0
	var currentInstruction1 int = 0

	timeout0 := make(chan bool)
	go func() {
		time.Sleep(deadlockTimeout * time.Second)
		timeout0 <- true
	}()
	timeout1 := make(chan bool)
	go func() {
		time.Sleep(deadlockTimeout * time.Second)
		timeout1 <- true
	}()

	channel0to1:= make(chan int,100000)
	channel1to0:= make(chan int,100000)

	channels:=map[int]chan int {0:channel0to1,1:channel1to0}

	go runProgram(0,instructions,registers0,currentInstruction0,channels,timeout0)
	go runProgram(1,instructions,registers1,currentInstruction1,channels,timeout1)

	time.Sleep(deadlockTimeout * time.Second)
}

func runProgram(id int, instructions []string ,registers map[string]int, currentInstruction int, channels map[int]chan int, timeout chan bool) {
	for currentInstruction>=0 && currentInstruction<len(instructions) {
		instruction:=instructions[currentInstruction]
		currentInstruction = RunInstructionPart2(id,instruction,registers,currentInstruction,channels,timeout)
	}
}



func RunInstructionPart2(id int,instruction string, registers map[string]int, currentInstruction int,channels map[int]chan int, timeout chan bool) int {
	arguments:=strings.Split(instruction," ")
	cmd:= arguments[0]
	var result int = currentInstruction + 1
	switch cmd {
	case "snd":
		x:=arguments[1]
		fmt.Println("[from",id, "]",cmd,x)
		val:=GetValue(x,registers)
		fmt.Println("[from",id, "] Sending value:",val,"on channel",id)
		channels[id] <- val
		var count int=GetValue(sndCounter,registers)
		registers[sndCounter]=count+1

	case "set":
		x:=arguments[1]
		y:=arguments[2]
		fmt.Println("[from",id, "]",cmd,x,y)
		val:=GetValue(y,registers)
		registers[x]=val
	case "add":
		x:=arguments[1]
		y:=arguments[2]
		fmt.Println("[from",id, "]",cmd,x,y)
		val1:=GetValue(x,registers)
		val2:=GetValue(y,registers)
		registers[x]=val1+val2
	case "mul":
		x:=arguments[1]
		y:=arguments[2]
		fmt.Println("[from",id, "]",cmd,x,y)
		val1:=GetValue(x,registers)
		val2:=GetValue(y,registers)
		registers[x]=val1*val2
	case "mod":
		x:=arguments[1]
		y:=arguments[2]
		fmt.Println("[from",id, "]",cmd,x,y)
		val1:=GetValue(x,registers)
		val2:=GetValue(y,registers)
		registers[x]=val1%val2
	case "rcv":
		x:=arguments[1]
		fmt.Println("[from",id, "]",cmd,x, "on channel ",1-id)
		c:=channels[1-id]
		select {
		case val := <-c:
			fmt.Println("[from", id, "] Receiving value:", val)
			registers[x] = val

		case <-timeout:
			fmt.Println("[from", id, "] TIMEOUT!!!")
			fmt.Println("register ",id,registers)
		}

	case "jgz":
		register1:=arguments[1]
		register2:=arguments[2]
		fmt.Println("[from",id, "]",cmd,register1,register2)
		val1:=GetValue(register1,registers)
		if val1 > 0 {
			val2:=GetValue(register2,registers)
			result=currentInstruction+val2
		}
	}
	return result
}

func RunInstruction(instruction string, registers map[string]int, currentInstruction int) (int,bool) {
	arguments:=strings.Split(instruction," ")
	cmd:= arguments[0]
	var result int = currentInstruction + 1
	var rcv bool = false
	switch cmd {
	case "snd":
		x:=arguments[1]
		fmt.Println(cmd,x)
		val:=GetValue(x,registers)
		registers[lastSoundPlayed]=val
		fmt.Println("Playing Sound",val)
	case "set":
		x:=arguments[1]
		y:=arguments[2]
		fmt.Println(cmd,x,y)
		val:=GetValue(y,registers)
		registers[x]=val
	case "add":
		x:=arguments[1]
		y:=arguments[2]
		fmt.Println(cmd,x,y)
		val1:=GetValue(x,registers)
		fmt.Println(val1)
		val2:=GetValue(y,registers)
		registers[x]=val1+val2
	case "mul":
		x:=arguments[1]
		y:=arguments[2]
		fmt.Println(cmd,x,y)
		val1:=GetValue(x,registers)
		val2:=GetValue(y,registers)
		registers[x]=val1*val2
	case "mod":
		x:=arguments[1]
		y:=arguments[2]
		fmt.Println(cmd,x,y)
		val1:=GetValue(x,registers)
		val2:=GetValue(y,registers)
		registers[x]=val1%val2
	case "rcv":
		register:=arguments[1]
		fmt.Println(cmd,register)
		val:=GetValue(register,registers)
		if val != 0 {
			frequencyRecovered:=GetValue(lastSoundPlayed,registers)
			fmt.Println("Recovered Last Sound Played frequency",frequencyRecovered)
			rcv=true
		}
	case "jgz":
		register1:=arguments[1]
		register2:=arguments[2]
		fmt.Println(cmd,register1,register2)
		val1:=GetValue(register1,registers)
		if val1 > 0 {
			val2:=GetValue(register2,registers)
			result=currentInstruction+val2
		}
	}
	return result,rcv
}

func GetValue(register string, registers map[string]int) int {
	var value int
	value,e:=strconv.Atoi(register)
	if e!= nil {
		v,present:=registers[register]
		if !present {
			value=0
			registers[register]=value
		} else {
			value=v
		}
	}
	return value
}