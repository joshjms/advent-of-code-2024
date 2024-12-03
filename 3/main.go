package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var s string

func parseInput(r *bytes.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		s += line
	}
}

func main() {
	filePath := "3/input.txt"

	b, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(b)
	parseInput(r)

	ans := 0

	// regexStr := `mul\(\d{1,3},\d{1,3}\)` // PART 1
	regexStr := `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindAllStringSubmatch(s, -1)

	do := true

	for _, match := range matches {
		m := match[0]
		fmt.Println(m)

		// PART 2
		if m == "do()" {
			do = true
			continue
		} else if m == "don't()" {
			do = false
			continue
		}
		// REMOVE THIS PART FOR PART 1

		if !do {
			continue
		}

		mulExp := m[4 : len(m)-1]
		fmt.Println(mulExp)

		intsStr := strings.Split(mulExp, ",")
		a, _ := strconv.Atoi(intsStr[0])
		b, _ := strconv.Atoi(intsStr[1])
		ans += a * b
	}

	fmt.Println(ans)
}
