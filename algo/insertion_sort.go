package algo

func InsertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > tmp {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = tmp

	}
	return arr
}
