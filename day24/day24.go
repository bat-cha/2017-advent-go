package day24

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Port struct {
	pins0 int
	pins1 int
	strength int
}

type Ports []Port

func (p Port) String() string {
	return fmt.Sprint(p.pins0, "/", p.pins1)
}

func (p Port) GetCompatibility(other Port) int {
	if p.pins1 == other.pins0 {
		return 0
	} else if p.pins1 == other.pins1 {
		return 1
	} else {
		return -1
	}
}

func (p *Port) Swap() {
	p.pins0,p.pins1=p.pins1,p.pins0
}

func (p Ports) CopyAndRemove(index int) Ports {
	result := Ports{}
	for i := 0; i < len(p); i++ {
		if i!=index {
			result = append(result, p[i])
		}
	}
	return result
}

func (p Ports) String() string {
	var result string
	for _, p := range p {
		result += p.String() + "--"
	}
	return result
}

func (p Ports) ComputeStrength() int {
	var result int
	for _, p := range p {
		result += p.strength
	}
	return result
}

var maxStrength int = 0
var maxLength int = 0
var maxStrengthLongest int = 0

func search(bridge Ports, remaining Ports) {
	port := bridge[len(bridge)-1]
	var noMoreMatch bool = true
	for i, p := range remaining {
		c:= port.GetCompatibility(p)
		if c != -1 {
			noMoreMatch = false
			if c==1 { p.Swap() }
			search(append(bridge,p), remaining.CopyAndRemove(i))
		}
	}
	if noMoreMatch {
		s := bridge.ComputeStrength()
		l := len(bridge)
		if s > maxStrength {
			maxStrength = s
		}
		if l==maxLength && s> maxStrengthLongest {
			maxStrengthLongest =s
		} else if l > maxLength {
			maxLength = l
			maxStrengthLongest = s
		}
	}

}

func Solve(filename string) int {
	zeroPorts, ports := initialize(filename)
	SolveParts(zeroPorts, ports)
	return maxStrength

}

func SolveParts(zeroPorts Ports, ports Ports) int {
	fmt.Println("bulding bridges from zero ports", zeroPorts, "with ports ", ports)
	for i, start := range zeroPorts {
		search(Ports{start},append(ports,zeroPorts.CopyAndRemove(i)...))
	}

	fmt.Println("Max Strengh", maxStrength)
	fmt.Println("Max Length", maxLength)
	fmt.Println("Max Strength Longest", maxStrengthLongest)

	return maxStrength

}

func initialize(filename string) (Ports, Ports) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()

	var ports Ports
	var zeroPorts Ports
	for e == nil {

		pins := strings.Split(string(s), "/")
		pins0, _ := strconv.Atoi(pins[0])
		pins1, _ := strconv.Atoi(pins[1])
		if pins0 == 0 || pins1 == 0 {
			zeroPorts = append(zeroPorts, Port{pins0, pins1, pins0+pins1})
		} else {
			ports = append(ports, Port{pins0, pins1, pins0+pins1})
		}

		s, _, e = r.ReadLine()
	}
	return zeroPorts, ports
}