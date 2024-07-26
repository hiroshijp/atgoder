package algo

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arg := []int{5, 3, 2, 4, 1}
	expected := []int{1, 2, 3, 4, 5}

	t.Run("BubbleSort", func(t *testing.T) {
		actual := BubbleSort(arg)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v but got %v", expected, actual)
		}
	})
}
