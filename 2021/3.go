package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	var lines []string
	lines = append(lines, line)
	text := strings.Split(line, "")
	var ray []string

	for _, s := range text {
		ray = append(ray, s)
	}
	for scanner.Scan() {
		line = scanner.Text()
		lines = append(lines, line)
		for i, s := range strings.Split(line, "") {
			ray[i] += s
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	part1(ray)
	println(part2(lines, true) * part2(lines, false))

}

func part1(ray []string) {
	var gamma, epsilon string
	for _, s := range ray {
		if strings.Count(s, "1") > strings.Count(s, "0") {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaDecimal, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDecimal, _ := strconv.ParseInt(epsilon, 2, 64)
	println(gammaDecimal * epsilonDecimal)
}

func part2(lines []string, isOnes bool) int64 {
	var searchString string

	for {
		var ones, zeroes = 0, 0
		for _, s := range lines {
			if strings.HasPrefix(s, searchString+"1") {
				ones += 1
			} else if strings.HasPrefix(s, searchString+"0") {
				zeroes += 1
			}
		}

		if ones >= zeroes && isOnes {
			searchString += "1"
		} else if ones < zeroes && isOnes {
			searchString += "0"
		} else if zeroes <= ones && !isOnes {
			searchString += "0"
		} else if zeroes > ones && !isOnes {
			searchString += "1"
		} else {
			println("allt är fukd")
		}
		if zeroes+ones == 2 {
			for _, s := range lines {
				if strings.HasPrefix(s, searchString) {
					searchInt, _ := strconv.ParseInt(s, 2, 64)
					return searchInt
				}
			}
		}
	}
}
