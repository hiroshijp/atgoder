package algo

import (
	"reflect"
	"testing"
)

func TestAsertionSort(t *testing.T) {
	args := []int{5, 3, 2, 4, 1}
	expected := []int{1, 2, 3, 4, 5}

	t.Run("AsertionSort", func(t *testing.T) {
		actrual := AsertionSort(args)
		if !reflect.DeepEqual(actrual, expected) {
			t.Errorf("expected %v but got %v", expected, actrual)
		}
	})
}
