package day08
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func Solve(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()

	registers:=map[string]int{}
	var highest int = 0

	for e == nil {
		line := string(s)
		elements := strings.Split(line, " if ")
		conditions := strings.Split(elements[1], " ")
		if Check(conditions,registers) {
			highest = Max(highest,ComputeInstructions(strings.Split(elements[0], " "), registers))
		}
		s, _, e = r.ReadLine()
	}
	var max int = 0
	for _,v := range registers {
		max=Max(max,v)
	}
	fmt.Println("the largest value in any register is",max)
	fmt.Println("the hightes value in any register during the process was ",highest)
}

func Check(conditions []string, registers map[string]int) bool {
	var check bool
	register:=conditions[0]
	registerValue:=getRegisterValue(register,registers)
	operand:=conditions[1]
	conditionValue,_:=strconv.Atoi(conditions[2])
	switch operand {
		case "<":
			check = registerValue < conditionValue
		case "<=":
			check = registerValue <= conditionValue
		case "==":
			check = registerValue == conditionValue
	    case "!=":
		    check = registerValue != conditionValue
		case ">=":
			check = registerValue >= conditionValue
		case ">":
			check = registerValue > conditionValue

	}

	//fmt.Println("conditions:",register,registerValue,operand,conditionValue,check)
	return check

}

func ComputeInstructions(actions []string, registers map[string]int) int {
	register:=actions[0]
	registerValue:=getRegisterValue(register,registers)
	operand:=actions[1]
	value,_:=strconv.Atoi(actions[2])
	var newValue int
	switch operand {
		case "inc":
			newValue=registerValue+value
		case "dec":
			newValue=registerValue-value
	}
	registers[register]=newValue
	//fmt.Println("actions:",register,registerValue,operand,value,newValue)
	return newValue

}

func getRegisterValue(register string, registers map[string]int) int {
	var registerValue int
	var present bool
	registerValue,present=registers[register]
	if !present {
		registerValue=0
	}
	return registerValue
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}