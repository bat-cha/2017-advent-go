package day14
import (
	"fmt"
	"bufio"
	"os"
	"../day10"
	"strconv"
	"strings"
)

type empty struct {

}

func Solve(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file= ",err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, _, e := r.ReadLine()
	for e == nil {
		var line string = string(s)
		SolvePart1(line)
		SolvePart2(line)
		s, _, e = r.ReadLine()
	}
}

func SolvePart1(key string) int {

	var result int = 0

	for row:=0;row<128;row++ {
		rowKey:=key+"-"+fmt.Sprint(row)
		//fmt.Println(rowKey)
		hash:=day10.KnotHash(rowKey)
		//fmt.Println(hash)
		bin:=HexToBin(hash)
		fmt.Println(bin)
		result+=strings.Count(bin,"1")
	}

	fmt.Println("Part 1 Result ",result)

	return result
}


func SolvePart2(key string) int {

	grid:=make([][]int,128)

	for row:=0;row<128;row++ {
		rowKey:=key+"-"+fmt.Sprint(row)
		hash:=day10.KnotHash(rowKey)
		bin:=HexToBin(hash)
		grid[row]=StrToSlice(bin)

	}
	var regions int = 0
	var oneFound bool = true
	for oneFound {
		oneFound=false
		indices:=map[string]empty{}
		for i := 0; i < 128; i++ {
			for j := 0; j < 128; j++ {
				if grid[i][j] == 1 {
					oneFound=true
					//remove all connected then increment region number
					RemoveConnected(i,j,grid,indices)
					for k,_:=range indices {
						i,j:=ToIndices(k)
						grid[i][j]=0
					}
					indices=map[string]empty{}
					regions++

				}

			}
		}


	}
	fmt.Println("Part 2  #regions ",regions)
	return regions
}

func HexToBin(hex string) string {
	var result=""
	for _,c:=range hex {
		ui, _ := strconv.ParseUint(string(c), 16, 32)
		result+=fmt.Sprintf("%04b", ui)
	}

	return result
}
func StrToSlice(binRow string) []int {
	result:=make([]int,len(binRow))
	for i,b:= range binRow {
		result[i],_=strconv.Atoi(string(b))
	}
	return result
}

func RemoveConnected(i int, j int, grid [][]int, indices map[string]empty) {
	indices[ToKey(i,j)]=empty{}
	var present bool
	_,present = indices[ToKey(i-1,j)]
	if !present && i>0 && grid[i-1][j]==1 {
		RemoveConnected(i-1,j,grid,indices)
	}
	_,present = indices[ToKey(i,j-1)]
	if !present && j>0 && grid[i][j-1]==1 {
		RemoveConnected(i,j-1,grid,indices)
	}
	_,present = indices[ToKey(i,j+1)]
	if !present && j<127 && (grid[i][j+1]==1) {
		RemoveConnected(i,j+1,grid,indices)
	}
	_,present = indices[ToKey(i+1,j)]
	if !present && i<127 && (grid[i+1][j]==1) {
		RemoveConnected(i+1,j,grid,indices)
	}

}

func ToKey(i int, j int) string {
	return fmt.Sprint(i,"-",j)
}

func ToIndices(key string) (int,int) {
	indices:=strings.Split(key,"-")
	i,_:=strconv.Atoi(indices[0])
	j,_:=strconv.Atoi(indices[1])
	return i,j
}