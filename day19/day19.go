package day19
import (
	"fmt"
	"bufio"
	"os"
	"regexp"
)

type position struct {
	row int
	column int
	value string
}

type direction struct {
	row int
	column int
}

var up direction = direction{-1,0}
var down direction = direction{1,0}
var left direction = direction{0,-1}
var right direction = direction{0,1}

var letter = regexp.MustCompile(`[A-Z]`)
type packet struct {
	direction *direction
	current *position
	letters []string
	steps int
}

func (p *packet) route (network map[int]map[int]string) {
	var newDirection *direction=nil
	switch {
	case p.current.value == "|" ||  p.current.value == "-" :
		newDirection = p.direction

	case p.current.value == "+":
		//find closest value to determine direction
		newDirection = p.current.LookAround(p.direction,network)

	case letter.MatchString(p.current.value):
		p.letters=append(p.letters,p.current.value)
		newDirection = p.direction

	}
	if newDirection != nil {
		p.steps++
		p.direction = newDirection
		p.current.Move(p.direction,network)
		p.route(network)
	} else {
		fmt.Println("Routing Finished ! Collected letters ",p.letters, " in ",p.steps,"steps")
	}


}

func (d *direction) Equals(other *direction) bool {
	return d.column==other.column && d.row==other.row
}

func (d *direction) GetOpposite() *direction {
	return &direction{-d.row,-d.column}
}

func (p *position) Move(d *direction, network map[int]map[int]string) {
	p.row = p.row + d.row
	p.column = p.column + d.column
	p.value = network[p.row][p.column]
}

func (p *position) LookAround(d *direction, network map[int]map[int]string) *direction {
	possible:=[]*direction{&left,&right,&up,&down}
	for _,try:=range possible {
		if !try.Equals(d) && !try.Equals(d.GetOpposite()){
			_,present:=network[p.row+try.row][p.column+try.column]
			if present {
				return try
			}
		}
	}
	return nil
}

func Solve(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()

	network :=map[int]map[int]string{}
	var rowIndex int = 0
	for e == nil {
		row:=GetRow(rowIndex, network)
		for i,c:=range []rune(string(s)) {
			char:=string(c)
			if char != " " {
				row[i]=char
			}
		}
		rowIndex++
		s, _, e = r.ReadLine()
	}
	//starting first value encountered on row 0
	var startRow int = 0
	var startCol int
	var startVal string
	for startCol, startVal = range network[startRow] { break }
	p:=packet{&down,&position{0,startCol,startVal},[]string{},0}

	p.route(network)


}


func SolvePart2(filename string) {

}

func GetRow(row int, grid map[int]map[int]string) map[int]string {

	result,present:= grid[row]
	if !present {
		result=map[int]string{}
		grid[row]=result
	}
	return result
}