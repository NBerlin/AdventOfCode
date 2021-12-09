package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	x := 0
	y := 0
	depth := 0

	file, err := os.Open("2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		row := strings.Split(scanner.Text(), " ")
		moveValue, _ := strconv.Atoi(row[1])
		switch row[0] {
		case "forward":
			x += moveValue
			depth += moveValue * y

		case "down":
			y += moveValue

		case "up":
			y -= moveValue
		}
	}
	println("Part 1 : ", x*y)
	println("part 2: ", x*depth)

}
