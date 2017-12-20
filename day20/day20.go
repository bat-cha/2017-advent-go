package day20

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

type coordinates struct {
	x float64
	y float64
	z float64
}

type particle struct {
	id           int
	position     coordinates
	velocity     coordinates
	acceleration coordinates
	destroyed 	 bool
}

func (p *particle) Update() {
	if !p.destroyed {
		p.velocity.x += p.acceleration.x
		p.velocity.y += p.acceleration.y
		p.velocity.z += p.acceleration.z
		p.position.x += p.velocity.x
		p.position.y += p.velocity.y
		p.position.z += p.velocity.z
	}
}

func (p *particle) CollideWith(other *particle) bool {
	return (p.position.x == other.position.x) && (p.position.y == other.position.y) && (p.position.z == other.position.z)
}

func (p *particle) ComputeDistance() float64 {
	return math.Abs(p.position.x) + math.Abs(p.position.y) + math.Abs(p.position.z)
}

func Solve(filename string) {
	SolvePart1(filename)
	SolvePart2(filename)
}

func SolvePart1(filename string) {

	buffer:=initialize(filename)
	var iterations int = 300
	for i := 0; i < iterations; i++ {
		for _, p := range buffer {
			p.Update()
		}
	}

	var min float64 = math.MaxFloat64
	var minId int = -1
	for id, p := range buffer {
		d := p.ComputeDistance()
		if d < min {
			min = d
			minId = id
		}
	}

	fmt.Println("particle", minId, "stays closer to zero, distance", min)

}

func SolvePart2(filename string) {
	buffer:=initialize(filename)
	var iterations int = 100
	var collisions int = 0
	for i := 0; i < iterations; i++ {
		for _, p := range buffer {
			p.Update()
		}
		for i, p1 := range buffer {
			if !p1.destroyed {
				for j:=i+1; j<len(buffer); j++ {
					p2:=buffer[j]
					if !p2.destroyed {
						if p1.CollideWith(p2) {
							p1.destroyed=true
							p2.destroyed=true
							collisions++
						}
					}
				}
			}
		}
	}
	var alive int = 0
	for _, p := range buffer {
		if !p.destroyed {
			alive++
		}

	}
	fmt.Println("particles left", alive, ", collisions #", collisions)
}

func initialize(filename string) []*particle {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()

	buffer := []*particle{}
	var particleId int = 0
	for e == nil {
		pva := strings.Split(string(s), ", ")
		p := pva[0][3:len(pva[0])-1]
		v := pva[1][3:len(pva[1])-1]
		a := pva[2][3:len(pva[2])-1]

		part := particle{particleId, *ParseCoordinates(p), *ParseCoordinates(v), *ParseCoordinates(a), false}
		buffer = append(buffer, &part)
		particleId++
		s, _, e = r.ReadLine()
	}
	return buffer
}

func ParseCoordinates(input string) *coordinates {

	values := strings.Split(input, ",")
	x, _ := strconv.Atoi(values[0])
	y, _ := strconv.Atoi(values[1])
	z, _ := strconv.Atoi(values[2])

	return &coordinates{float64(x), float64(y), float64(z)}

}


