package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseReader(r io.Reader) ([][]int, error) {
	var result [][]int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Skip empty lines if needed
			continue
		}

		// Split the line into strings by spaces
		parts := strings.Fields(line)

		// Convert the parts to integers
		var nums []int
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("error parsing number %q: %v", part, err)
			}
			nums = append(nums, num)
		}

		// Append the line of numbers to the result
		result = append(result, nums)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %v", err)
	}

	return result, nil
}

func checkLine(line []int) int {
	pos := 0
	neg := 0

	for i := 0; i < len(line)-1; i++ {
		d := line[i] - line[i+1]
		if d > 0 {
			pos = 1
		} else if d < 0 {
			neg = 1
		} else {
			return i
		}

		if pos == 1 && neg == 1 {
			return i
		}

		if math.Abs(float64(d)) > 3 {
			return i
		}
	}

	return -1
}

func main() {
	filePath := "2/input.txt"

	b, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(b)

	data, err := parseReader(r)
	if err != nil {
		panic(err)
	}

	ans := 0

	for _, line := range data {
		err := checkLine(line)

		if err != -1 {
			for i := 0; i < len(line); i++ {
				tmpLine := make([]int, 0)
				tmpLine = append(tmpLine, line[:i]...)
				tmpLine = append(tmpLine, line[i+1:]...)

				e := checkLine(tmpLine)
				if e == -1 {
					ans++
					break
				}
			}
		} else {
			ans++
		}

	}

	fmt.Println(ans)
}
