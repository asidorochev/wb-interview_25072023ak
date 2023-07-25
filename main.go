package main

import (
	"fmt"
	"os"
)

func main() {
	var res []string

	defer func(res []string) {
		fmt.Println("result =", res)
	}(res)

	length := len("Привет")

	data := make([]int, 0, length)
	for i := 0; i < length; i++ {
		data[i] = i
	}

	addSumValue(data)

	var multiplied map[int]int
	for i, v := range data {
		multiplied[i] = v * 2
	}

	ch := make(chan int)

	for _, v := range multiplied {
		ch <- v
	}

	var i int

	go func() {
		for v := range ch {
			i++
			res = append(res, fmt.Sprintf("%d:%d", i, v))
		}

		fmt.Println("job 1 is finished")
	}()

	go func() {
		for v := range ch {
			i++
			res = append(res, fmt.Sprintf("%d:%d", i, v))
		}

		fmt.Println("job 2 is finished")
	}()

	os.Exit(0)
}

func addSumValue(data []int) {
	var sum int
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}

	data = append(data, sum)
}
