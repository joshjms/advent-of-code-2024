package main

import (
	"bytes"
	"fmt"
	"os"
)

type input struct {
	a []int64
	b []int64
}

func main() {
	filePath := "1/input.txt"

	b, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(b)

	var n = 1000

	input := input{
		a: make([]int64, n),
		b: make([]int64, n),
	}

	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d %d\n", &input.a[i], &input.b[i])
	}

	// d := 0

	// sort.Slice(input.a, func(i, j int) bool {
	// 	return input.a[i] < input.a[j]
	// })

	// sort.Slice(input.b, func(i, j int) bool {
	// 	return input.b[i] < input.b[j]
	// })

	// for i := 0; i < n; i++ {
	// 	d += int(math.Abs(float64(input.a[i] - input.b[i])))
	// }

	// fmt.Println(d)

	mp := make(map[int64]int64)

	for i := 0; i < n; i++ {
		mp[input.b[i]]++
	}

	s := int64(0)

	for i := 0; i < n; i++ {
		s += input.a[i] * mp[input.a[i]]
	}

	fmt.Println(s)

}
