package day22
import (
	"fmt"
	"bufio"
	"os"
)

type node struct {
	row int
	column int
}

type direction struct {
	row int
	column int
}

func (d *direction) Equals(other *direction) bool {
	return d.column==other.column && d.row==other.row
}

func (d *direction) GetOpposite() *direction {
	return &direction{-d.row,-d.column}
}

var up direction = direction{-1,0}
var down direction = direction{1,0}
var left direction = direction{0,-1}
var right direction = direction{0,1}

type virus struct {
	direction *direction
	current *node
	bursts int
	infections int
	cleaned int
	weakened int
	flagged int
}

type status struct {
	infected bool
	weakened bool
	flagged bool
}

func (p *node) Move(d *direction) {
	p.row = p.row + d.row
	p.column = p.column + d.column
}

func (v *virus) Burst(infectedNodes map[int]map[int]*status) {
	infected:=IsInfectedPart1(v.current,infectedNodes)
	if infected {
		v.direction=v.GetTurnRightDirection()
		Clean(v.current,infectedNodes)
		v.cleaned++
	} else {
		v.direction=v.GetTurnLeftDirection()
		Infect(v.current,infectedNodes)
		v.infections++
	}
	v.current.Move(v.direction)
	v.bursts++
}

func (v *virus) Burst2(infectedNodes map[int]map[int]*status) {
	status:=GetStatus(v.current,infectedNodes)
	switch {
	case status.weakened:
		status.weakened=false
		status.infected=true
		v.infections++

	case status.infected:
		v.direction=v.GetTurnRightDirection()
		status.infected=false
		status.flagged=true
		v.flagged++

	case status.flagged:
		v.direction=v.direction.GetOpposite()
		Clean(v.current,infectedNodes)
		v.cleaned++

	default:
		v.direction=v.GetTurnLeftDirection()
		status.weakened=true
		v.weakened++
	}

	v.current.Move(v.direction)
	v.bursts++
}

func (v *virus) GetTurnLeftDirection() *direction {
	var result *direction = nil
	switch {
	case v.direction.Equals(&up):
		result = &left

	case v.direction.Equals(&down):
		result = &right

	case v.direction.Equals(&left):
		result = &down

	case v.direction.Equals(&right):
		result = &up
	}
	return result
}

func (v *virus) GetTurnRightDirection() *direction {
	var result *direction = nil
	switch {
	case v.direction.Equals(&up):
		result = &right

	case v.direction.Equals(&down):
		result = &left

	case v.direction.Equals(&left):
		result = &up

	case v.direction.Equals(&right):
		result = &down
	}
	return result
}

func Solve(filename string) {
	virus,infectedNodes:=initialize(filename)
	SolvePart1(virus,10000,infectedNodes)
	virus,infectedNodes=initialize(filename)
	SolvePart2(virus,10000000,infectedNodes)

}

func SolvePart1(virus *virus, bursts int, infectedNodes map[int]map[int]*status) int {

	for burst:=0;burst<bursts;burst++ {
		virus.Burst(infectedNodes)
	}
	fmt.Println("#burst:",virus.bursts, "#infections:",virus.infections,"#cleaned:",virus.cleaned)
	return virus.infections


}

func SolvePart2(virus *virus, bursts int, infectedNodes map[int]map[int]*status) int {

	for burst:=0;burst<bursts;burst++ {
		virus.Burst2(infectedNodes)
	}
	fmt.Println("#burst:",virus.bursts, "#infections:",virus.infections,"#cleaned:",virus.cleaned, "#weakened:",virus.weakened,"#flagged:",virus.flagged)
	return virus.infections


}

func initialize(filename string) (*virus,map[int]map[int]*status) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()

	infectedNodes :=map[int]map[int]*status{}
	var rowIndex int = 0
	for e == nil {
		row:=GetRow(rowIndex, infectedNodes)
		for i,c:=range []rune(string(s)) {
			char:=string(c)
			if char == "#" {
				row[i]=&status{true,false,false}
			}
		}
		rowIndex++
		s, _, e = r.ReadLine()
	}
	//starting in middle of the grid
	startRow:=rowIndex/2
	startCol:=rowIndex/2
	virus:=virus{&up,&node{startRow,startCol},0,0,0,0,0}

	return &virus,infectedNodes



}

func GetRow(row int, infectedNodes map[int]map[int]*status) map[int]*status {

	result,present:= infectedNodes[row]
	if !present {
		result=map[int]*status{}
		infectedNodes[row]=result
	}
	return result
}

func IsInfectedPart1(node *node,infectedNodes map[int]map[int]*status) bool{
	_,infected:= infectedNodes[node.row][node.column]
	return infected
}

func Clean(node *node,infectedNodes map[int]map[int]*status) {
	delete(infectedNodes[node.row],node.column)
}

func Infect(node *node,infectedNodes map[int]map[int]*status) {
	GetRow(node.row,infectedNodes)[node.column]=&status{true,false,false}
}


func GetStatus(node *node,infectedNodes map[int]map[int]*status) *status {
	var nodeStatus *status
	nodeStatus,present:=GetRow(node.row,infectedNodes)[node.column]
	if !present {
		nodeStatus=&status{false,false,false}
		infectedNodes[node.row][node.column]=nodeStatus
	}
	return nodeStatus
}