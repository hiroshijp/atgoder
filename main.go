package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)

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

func nextIntWithSpace() []int {
	sc.Scan()
	list := make([]int, 10)
	for _, e := range strings.Split(sc.Text(), " ") {
		n, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}
		list = append(list, n)
	}
	return list
}

func nextStringWithSpace() []string {
	sc.Scan()
	list := strings.Split(sc.Text(), " ")
	return list
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
