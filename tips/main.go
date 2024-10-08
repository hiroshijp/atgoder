package main

import (
	"fmt"
	"strconv"
)

func main() {
	grid := [][]string{
		{"0", "1", "0", "0"},
		{"1", "1", "1", "0"},
		{"0", "0", "0", "1"},
		{"0", "1", "1", "0"},
	}

	grid = tipGridGravity(grid, "0")

	for _, el := range grid {
		fmt.Println(el)
	}
}

func tipString() {
	en := "abcde"
	b := []byte(en)
	b[0] = 'k'
	fmt.Println(string(b))

	ja := "あいうえお"
	r := []rune(ja)
	r[0] = 'か'
	fmt.Println(string(r))
	fmt.Println(string(r[:3]))
}

func tipSlice() {
	arr := []int{1, 2, 3, 4, 5, 6}
	arr1 := arr[0:3]
	arr2 := arr[3:len(arr)]
	fmt.Println(arr1, arr2)
}

func tipToString() {
	n := 10
	b := strconv.FormatInt(int64(n), 2)
	fmt.Println(b)
}

func tipMap() {
	m := make(map[int]int)
	for _, e := range []int{1, 1, 2, 3, 3} {
		if v, ok := m[e]; !ok {
			m[e] = 1
		} else {
			m[e] = v + 1
		}
	}
	fmt.Println(m)
}

func tipSuretsu() {
	arr := []int{1, 2, 3, 4, 5}
	odd := len(arr)

	sum := 0
	if odd%2 == 1 {
		sum = ((odd + 1) / 2) * odd
	} else {
		sum = (odd / 2) * (odd + 1)
	}
	fmt.Println(sum)
}

func tipGridGravity(grid [][]string, space string) [][]string {
	row := len(grid)
	col := len(grid[0])

	for i := 0; i < col; i++ {
		tmp := []string{}
		for j := 0; j < row; j++ {
			if grid[row-1-j][i] != space {
				tmp = append(tmp, grid[row-1-j][i])
			}
		}
		for j := 0; j < row; j++ {
			if j < len(tmp) {
				grid[row-1-j][i] = tmp[j]
			} else {
				grid[row-1-j][i] = space
			}
		}
	}
	return grid
}
