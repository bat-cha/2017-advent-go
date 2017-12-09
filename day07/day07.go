package day07
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	//"sort"
)

type Program struct {
	name string
	weight int
	holdingPrograms []*Program
	bottomProgram *Program
	weightFix int
}

func (p *Program) String() string {
	return p.name
}

func (p *Program) ComputeTotalWeight() int {
	var totalWeight int = p.weight
	for _,h:= range p.holdingPrograms {
		totalWeight+=h.ComputeTotalWeight()
	}
	return totalWeight
}

func (p *Program) CheckBalance() bool {
	l := len(p.holdingPrograms)
	var result bool = true
	if l > 0 {

		holdingBalance := make([]int, l)
		holdingBalance[0]= p.holdingPrograms[0].ComputeTotalWeight()
		var min int = holdingBalance[0]
		for i, h := range p.holdingPrograms {
			holdingBalance[i] = h.ComputeTotalWeight()
			min = Min(min, holdingBalance[i])
		}
		for i, h := range p.holdingPrograms {
			holdingBalance[i] -=min
			if holdingBalance[i]>0 {
				if h.CheckBalance() {
					h.weightFix = - holdingBalance[i]
					result = false
				}
			}
		}
	}
	return result

}

func Solve(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()

	programs:=map[string]*Program{}

	for e == nil {
		var line string = string(s)
		var elements = strings.Split(line,"-> ")
		properties := strings.Split(elements[0]," ")
		name := properties[0]
		weight,_ := strconv.Atoi(strings.TrimPrefix(strings.TrimSuffix(properties[1],")"),"("))
		p,present := programs[name]
		if !present {
			p = &Program{name,0,nil,nil, 0}
			programs[name]=p
		}
		p.weight = weight

		if len(elements)>1 {

			holdingProgramsNames := strings.Split(elements[1],", ")
			holdingPrograms:=make([]*Program,len(holdingProgramsNames))
			for i, name := range holdingProgramsNames {
				h,present := programs[name]
				if present {
					h.bottomProgram=p
				} else {
					h = &Program{name,0,nil,p, 0}
					programs[name]=h
				}
				holdingPrograms[i]=h
			}
			p.holdingPrograms = holdingPrograms
		}

		s, _, e = r.ReadLine()
	}
	var bottom string
	var wrongWeight string
	var weightFixed int = 42000000
	for _, p := range programs {
		p.CheckBalance()
	}
	for _, p := range programs {
		if p.bottomProgram == nil {
			bottom = p.name
			fmt.Println(p.name," ",p.weight, " ",p.bottomProgram, " ", p.weightFix)
		}
		if p.weightFix != 0 {
			fmt.Println(p.name," ",p.weight, " ",p.bottomProgram, " ", p.weightFix)
			if p.weight < weightFixed {
				fmt.Println(p.name)
				wrongWeight = p.name
				weightFixed = p.weight + p.weightFix
			}


			fmt.Println(p.name," ",p.weight, " ",p.bottomProgram, " ", p.weightFix)
		}
	}
	fmt.Println("Bottom Program ", bottom)
	fmt.Println("Wrong weight Program ", wrongWeight, " correct weight", weightFixed)
}


func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}