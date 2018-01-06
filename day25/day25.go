package day25

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

var START string = "Begin in state "
var STEPS string = "Perform a diagnostic checksum after "
var STATE_START string = "In state "
var CHECK_START string = "  If the current value is "
var WRITE string = "    - Write the value "
var MOVE string = "    - Move one slot to the "
var CONTINUE string = "    - Continue with state "

type State struct {
	name  string
	write map[int]int
	move  map[int]int
	next  map[int]string
}

func (s State) String() string {
	return fmt.Sprint(s.name, " w", s.write, " m", s.move, " n", s.next)
}

type TuringMachine struct {
	tape          map[int]int
	cursor        int
	state         string
	checksumSteps int
	states        map[string]State
}

func (m TuringMachine) String() string {
	return fmt.Sprint(m.tape, " ", m.cursor, " ", m.state, " ", m.checksumSteps, " ", m.states)
}

func (m TuringMachine) ComputeCheckSum() int {
	var result int = 0
	for _,v:=range m.tape {
		result+=v
	}
	return result
}

func (m TuringMachine) Run() {
	var step int = 0
	for step < m.checksumSteps {
		s:=m.states[m.state]
		write:=s.write[m.tape[m.cursor]]
		move:=s.move[m.tape[m.cursor]]
		next:=s.next[m.tape[m.cursor]]
		m.tape[m.cursor]=write
		m.cursor+=move
		m.state=next
		step++
	}

}

func Solve(filename string) int {
	machine := initialize(filename)
	machine.Run()
	cs:=machine.ComputeCheckSum()
	fmt.Println("CheckSum",cs)
	return cs

}

func initialize(filename string) TuringMachine {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()

	var start string
	var checksumSteps int
	states := map[string]State{}
	var parsingState *State
	var parsingCursor int

	for e == nil {

		line := string(s)
		switch {
		case strings.HasPrefix(line, START):
			start = line[len(START):len(START)+1]

		case strings.HasPrefix(line, STEPS):
			checksumSteps, _ = strconv.Atoi(strings.Split(line[len(STEPS):len(line)-1], " ")[0])

		case strings.HasPrefix(line, STATE_START):
			if parsingState != nil {
				states[parsingState.name] = *parsingState
			}
			name := line[len(STATE_START):len(STATE_START)+1]
			parsingState = &State{name, map[int]int{}, map[int]int{}, map[int]string{}}

		case strings.HasPrefix(line, CHECK_START):
			parsingCursor, _ = strconv.Atoi(line[len(CHECK_START):len(CHECK_START)+1])

		case strings.HasPrefix(line, WRITE):
			value,_:=strconv.Atoi(line[len(WRITE):len(WRITE)+1])
			parsingState.write[parsingCursor]=value

		case strings.HasPrefix(line, MOVE):
			move:=line[len(MOVE):len(line)-1]
			switch move {
			case "right":
				parsingState.move[parsingCursor]=1
			case "left":
				parsingState.move[parsingCursor]=-1
			}

		case strings.HasPrefix(line,CONTINUE):
			parsingState.next[parsingCursor]=line[len(CONTINUE):len(CONTINUE)+1]

		}

		s, _, e = r.ReadLine()
	}
	if parsingState != nil {
		states[parsingState.name] = *parsingState
	}

	m := TuringMachine{map[int]int{0:0}, 0, start, checksumSteps, states}
	fmt.Println(m)
	return m
}
