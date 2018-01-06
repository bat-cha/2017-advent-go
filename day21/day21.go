package day21

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"log"
	"math"
)


func Solve(filename string) {
	SolvePart(filename,5)
	SolvePart(filename, 18)
}

func SolvePart(filename string, iterations int) int {
	rules:=initialize(filename)

	fractal := ".#./..#/###"
	var newFractal []string
	for it:=0;it<iterations;it++ {
		newFractal = nil
		squares:=Divide2Squares(fractal)

		for _,square:=range squares {
			newFractal=append(newFractal,GetNew(square,rules))
		}
		fractal = Merge2Square(newFractal)

	}
	result:=strings.Count(fractal,"#")
	fmt.Println(result,"pixels ON")
	return result

}

func GetNew(search string, rules map[string]string) string {
	var present bool = false
	var input string = search
	var output string
	for _,input:=range GetPatterns(search) {
		output,present = rules[input]
		if present {
			break
		}

	}
	if !present {
		log.Fatal("not found ",input)
	}
	return output
}

func Divide2Squares(input string) []string {
	var result []string = nil
	rows:=strings.Split(input,"/")
	if len(rows)%2 == 0 {
		for i:=0;i<len(rows)-1;i=i+2 {
			for j:=0;j<len(rows)-1;j=j+2 {
				square:=string(rows[i][j])+string(rows[i][j+1])+"/"+string(rows[i+1][j])+string(rows[i+1][j+1])
				result=append(result,square)
			}
		}

	} else {
		if len(rows)%3 == 0 {
			for i:=0;i<len(rows)-2;i=i+3 {
				for j:=0;j<len(rows)-2;j=j+3 {
					square:=string(rows[i][j])+string(rows[i][j+1])+string(rows[i][j+2])+"/"+string(rows[i+1][j])+string(rows[i+1][j+1])+string(rows[i+1][j+2])+"/"+string(rows[i+2][j])+string(rows[i+2][j+1])+string(rows[i+2][j+2])
					result=append(result,square)
				}
			}
		}
	}

	return result
}

func Merge2Square(inputs []string) string {
	numSquares := len(inputs)
	numRows:=int(math.Sqrt(float64(numSquares)))
	if numSquares >1 {
		var rows []string
		squareSize := len(strings.Split(inputs[0], "/"))

		for square := 0; square < numSquares-(numRows-1); square = square + numRows {
			var rowElement [][]string
			for r:=0;r<numRows;r++ {
				rowElement =append(rowElement,strings.Split(inputs[square+r], "/"))
			}
			for i := 0; i < squareSize; i++ {

				var row string
				for r:=0;r<numRows;r++ {
					row = row+rowElement[r][i]
				}
				rows = append(rows, row)
			}
		}

		result := strings.Join(rows, "/")
		return result
	} else {
		return inputs[0]
	}


}

func GetPatterns(input string) []string {
	var patterns []string = []string{input}
	rows:=strings.Split(input,"/")
	if len(rows)%2 == 0 {
		rot1:=string(rows[0][1])+string(rows[1][1])+"/"+string(rows[0][0])+string(rows[1][0])
		rot2:=string(rows[1][1])+string(rows[1][0])+"/"+string(rows[0][1])+string(rows[0][0])
		rot3:=string(rows[1][0])+string(rows[0][0])+"/"+string(rows[1][1])+string(rows[0][1])
		flip1:=string(rows[0][1])+string(rows[0][0])+"/"+string(rows[1][1])+string(rows[1][0])
		flip2:=string(rows[1][0])+string(rows[1][1])+"/"+string(rows[0][0])+string(rows[0][1])
		flip3:=string(rows[0][0])+string(rows[1][0])+"/"+string(rows[0][1])+string(rows[1][1])
		flip4:=string(rows[1][1])+string(rows[0][1])+"/"+string(rows[1][0])+string(rows[0][0])
		patterns=append(patterns, rot1, rot2, rot3, flip1, flip2, flip3, flip4)

	}else {
		if len(rows)%3 == 0 {
			rot1:=string(rows[0][2])+string(rows[1][2])+string(rows[2][2])+"/"+string(rows[0][1])+string(rows[1][1])+string(rows[2][1])+"/"+string(rows[0][0])+string(rows[1][0])+string(rows[2][0])
			rot2:=string(rows[2][2])+string(rows[2][1])+string(rows[2][0])+"/"+string(rows[1][2])+string(rows[1][1])+string(rows[1][0])+"/"+string(rows[0][2])+string(rows[0][1])+string(rows[0][0])
			rot3:=string(rows[2][0])+string(rows[1][0])+string(rows[0][0])+"/"+string(rows[2][1])+string(rows[1][1])+string(rows[0][1])+"/"+string(rows[2][2])+string(rows[1][2])+string(rows[0][2])
			flip1:=string(rows[0][2])+string(rows[0][1])+string(rows[0][0])+"/"+string(rows[1][2])+string(rows[1][1])+string(rows[1][0])+"/"+string(rows[2][2])+string(rows[2][1])+string(rows[2][0])
			flip2:=string(rows[2][0])+string(rows[2][1])+string(rows[2][2])+"/"+string(rows[1][0])+string(rows[1][1])+string(rows[1][2])+"/"+string(rows[0][0])+string(rows[0][1])+string(rows[0][2])
			flip3:=string(rows[0][0])+string(rows[1][0])+string(rows[2][0])+"/"+string(rows[0][1])+string(rows[1][1])+string(rows[2][1])+"/"+string(rows[0][2])+string(rows[1][2])+string(rows[2][2])
			flip4:=string(rows[2][2])+string(rows[1][2])+string(rows[0][2])+"/"+string(rows[2][1])+string(rows[1][1])+string(rows[0][1])+"/"+string(rows[2][0])+string(rows[1][0])+string(rows[0][0])
			patterns=append(patterns, rot1, rot2, rot3, flip1, flip2,flip3,flip4)

		}
	}
	return patterns
}

func initialize(filename string) map[string]string {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()

	rules := map[string]string{}

	for e == nil {
		io := strings.Split(string(s), " => ")
		rules[io[0]]=io[1]
		s, _, e = r.ReadLine()
	}
	return rules
}
