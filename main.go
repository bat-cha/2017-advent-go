package main

import (
	"./day01"
	"./day02"
	"./day03"
	"./day04"
	"./day05"
	"./day06"
	"./day07"
	"./day08"

	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
	"time"
	"log"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Hello, which Advent of Code 2017 day would you like to solve ?: ")
	text, e := reader.ReadString('\n')
	if e!= nil {
		fmt.Print("Error reading your input")
		os.Exit(1)
	}
	choice,e := strconv.Atoi(strings.TrimSuffix(text,"\n"))
	if e!= nil {
		fmt.Print("Wrong format, please input integer ",e)
		os.Exit(1)
	}

	instance:=fmt.Sprintf("solving puzzle %02d",choice)
	fmt.Println(instance)

	defer timeTrack(time.Now(), instance)

	var filename = fmt.Sprintf("day%02d/input.txt",choice)

	switch choice {

	case 1:
		day01.Solve(filename)
	case 2:
		day02.Solve(filename)
	case 3:
		day03.Solve(filename)
	case 4:
		day04.Solve(filename)
	case 5:
		day05.Solve(filename)
	case 6:
		day06.Solve(filename)
	case 7:
		day07.Solve(filename)
	case 8:
		day08.Solve(filename)

	default:
		fmt.Println("unknown puzzle")
		os.Exit(1)

	}

}


func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}