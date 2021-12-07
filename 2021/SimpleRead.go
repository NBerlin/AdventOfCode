package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// if parsing ints
	for scanner.Scan() {
		temp, _ := strconv.Atoi(scanner.Text())
		values = append(values, temp)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
