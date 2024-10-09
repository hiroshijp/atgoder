package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	step := nextString()
	switch step {
	case "1":
		step1()
	case "2":
		step2()
	case "3":
	case "4":
	}
}

func step1() {
	M := nextInt()

	type food struct {
		D int
		S int
		P int
	}
	foods := make(map[int]*food, M)
	for i := 0; i < M; i++ {
		d := nextInt()
		foods[d] = &food{d, nextInt(), nextInt()}
	}

	order := [][]int{}
	for {
		if nextString() != "order" {
			break
		}
		order = append(order, []int{nextInt(), nextInt(), nextInt()})
	}

	for _, el := range order {
		t, d, n := el[0], el[1], el[2]
		food := foods[d]
		if n <= food.S {
			for j := 0; j < n; j++ {
				food.S--
				fmt.Printf("received order %d %d\n", t, d)
			}
		} else {
			fmt.Printf("sold out %d\n", t)
		}
	}
}

func step2() {
	M := nextInt()
	K := nextInt()
	fmt.Println(K)

	type food struct {
		D int
		S int
		P int
	}
	foods := make(map[int]*food, M)
	for i := 0; i < M; i++ {
		d := nextInt()
		foods[d] = &food{d, nextInt(), nextInt()}
	}

	status := [][]int{}
	for {
		s := nextString()
		if s == "received" {
			nextString()
			status = append(status, []int{0, nextInt(), nextInt()})
		} else if s == "complete" {
			status = append(status, []int{1, nextInt()})
		} else {
			break
		}
	}
	fmt.Println(status)

}

func nextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func nextIntMatrix(row int, col int) [][]int {
	matrix := make([][]int, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]int, col)
		for j := 0; j < col; j++ {
			matrix[i][j] = nextInt()
		}
	}
	return matrix
}

func arrToCounetr(arr []int) map[int]int {
	m := make(map[int]int)
	for _, el := range arr {
		if v, ok := m[el]; !ok {
			m[el] = 1
		} else {
			m[el] = v + 1
		}
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func DebugIntMatrix(m [][]int) {
	for _, v := range m {
		fmt.Println(v)
	}
	fmt.Println("--------------------------------")
}

func DebugStringMatrix(m [][]string) {
	for _, v := range m {
		fmt.Println(v)
	}
	fmt.Println("--------------------------------")
}
