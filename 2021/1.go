package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func sum(test []int) int {
	sum := 0
	for _, v := range test {
		sum += v
	}
	return sum
}

func main() {
	file, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	counter := 0
	values := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		temp, _ := strconv.Atoi(scanner.Text())
		values = append(values, temp)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i, _ := range values {
		if sum(values[i:i+3]) < sum(values[i+1:i+4]) {
			counter++
		}
	}

	print(counter)
}
