package day11
import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
	"io/ioutil"
	"log"
	"strings"
)

var n = mat.NewVecDense(3,[]float64{0,1,-1})
var ne = mat.NewVecDense(3,[]float64{1,0,-1})
var nw = mat.NewVecDense(3,[]float64{-1,1,0})
var s = mat.NewVecDense(3,[]float64{0,-1,1})
var se = mat.NewVecDense(3,[]float64{1,-1,0})
var sw = mat.NewVecDense(3,[]float64{-1,0,1})

func Solve(filename string) {

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	} else {
		moves:=strings.Split(string(buf),",")
		distance,maxDistance:= SolvePuzzle(moves)
		fmt.Println("part 1 distance: ",distance)
		fmt.Println("part 2 max distance: ",maxDistance)
	}


}

func SolvePuzzle(moves []string) (int,int) {
	var result int = 0
	var maxDistance int = 0
	var position = mat.NewVecDense(3,[]float64{0,0,0})
	for _,move:= range moves {
		position.AddVec(position,getMove(move))
		maxDistance=Max(maxDistance,getDistance(position))

	}
	fmt.Println("final position in grid: ",position.At(0,0), position.At(1,0), position.At(2,0))
	result=getDistance(position)
	return result,maxDistance

}

func getMove(move string) *mat.VecDense {
	switch move {
		case "n": return n
		case "ne": return ne
		case "nw": return nw
		case "s": return s
		case "se": return se
		case "sw": return sw
	}
	return n
}

func getDistance(position *mat.VecDense) int {
	return int(math.Max(math.Max(math.Abs(position.At(0,0)),math.Abs(position.At(1,0))),math.Abs(position.At(2,0))))
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

