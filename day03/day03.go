package day03
import (
	"fmt"
	"strconv"
	"gonum.org/v1/gonum/mat"
	"math"
	"io/ioutil"
	"log"
)

var right = mat.NewVecDense(2,[]float64{0,1})
var left = mat.NewVecDense(2,[]float64{0,-1})
var up = mat.NewVecDense(2,[]float64{1,0})
var down = mat.NewVecDense(2,[]float64{-1,0})

type empty struct {

}

type content struct {
	value float64

}

func Solve(filename string) {

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	} else {

		var inputSquare int
		inputSquare, _ = strconv.Atoi(string(buf))
		fmt.Println("destination in spiral: ", inputSquare)
		//part1(inputSquare)
		//part1(5)
		part2(inputSquare)
		//part2(7)
	}


}

func part1(destination int) {
	//start from 1 [0,0]
	var present empty
	visited := map[string]empty {}
	var position = mat.NewVecDense(2,[]float64{0,0})
	visited[VecToKey(position)]=present

	var tryPosition = mat.NewVecDense(2,[]float64{0,0})
	var lastMove = mat.NewVecDense(2,[]float64{0,0})

	var occupied bool

	var count int = 1
	for count < destination {

		var move = NextMove(position,lastMove)
		tryPosition.AddVec(position,move)
		_, occupied = visited[VecToKey(tryPosition)]
		if (occupied) {
			position.AddVec(position,lastMove)
		} else {
			position.CopyVec(tryPosition)
			lastMove.CopyVec(move)

		}
		visited[VecToKey(position)]=present
		count++
	}
	fmt.Println("position in spiral: ",position.At(0,0), position.At(1,0))
	fmt.Println("#steps: ",math.Abs(position.At(0,0)) + math.Abs(position.At(1,0)))

}

func part2(destination int) {
	//start from 1 [0,0]
	square := content{float64(1)}
	visited := map[string]content {}
	var position = mat.NewVecDense(2,[]float64{0,0})
	visited[VecToKey(position)]=square

	var tryPosition = mat.NewVecDense(2,[]float64{0,0})
	var lastMove = mat.NewVecDense(2,[]float64{0,0})


	var occupied bool

	var count int = 1
	for count < destination && int(square.value) < destination {

		var move = NextMove(position,lastMove)
		tryPosition.AddVec(position,move)
		_, occupied = visited[VecToKey(tryPosition)]
		if (occupied) {
			position.AddVec(position,lastMove)
		} else {
			position.CopyVec(tryPosition)
			lastMove.CopyVec(move)

		}
		square = computeValue(position,visited)
		fmt.Println("putting ",square.value," into ", VecToKey(position))
		visited[VecToKey(position)]=square
		count++
	}

	fmt.Println("position in spiral: ",position.At(0,0), position.At(1,0))
	fmt.Println("value: ",square.value)

}

func NextMove(position *mat.VecDense, lastMove *mat.VecDense) *mat.VecDense {
	if Equal(lastMove, right) {
		return up
	}
	if Equal(lastMove, up) {
		return left
	}
	if Equal(lastMove, left) {
		return down
	}
	if Equal(lastMove, down) {
		return right
	}
	return right
}

func computeValue(position *mat.VecDense, visited map[string]content ) content {
	total := float64(0)
	//visit all neighbors, add content to total if any
	//fmt.Println("finding neighbors of ",VecToKey(position))
	var neighbor = mat.NewVecDense(2,[]float64{0,0})
	neighbor.CopyVec(position)
	path := [8]*mat.VecDense{left,up,right,right,down,down,left,left}
	for i := 0; i<8; i++ {
		neighbor.AddVec(neighbor,path[i])
		//fmt.Println(VecToKey(neighbor))
		square, occupied := visited[VecToKey(neighbor)]
		if (occupied) {
			//fmt.Println("found neighbor at ",VecToKey(neighbor)," value", square.value)
			total += square.value
		}
	}
	return content{total}

}

func Equal(v1 *mat.VecDense, v2 *mat.VecDense) bool {
	return v1.At(0,0) == v2.At(0,0) && v1.At(1,0) == v2.At(1,0)
}

func VecToKey(v1 *mat.VecDense) string {
	return strconv.FormatFloat(v1.At(0,0),'f', 0, 64) + strconv.FormatFloat(v1.At(1,0),'f', 0, 64)
}