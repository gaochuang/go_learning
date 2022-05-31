package arraytest_test

import "testing"

func TestArrayInit(t *testing.T) {
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{4, 5, 6, 7, 8}
	t.Log(arr1[1], arr1[2])
	t.Log(arr2)
}

func TestArratTravel(t *testing.T) {
	arr5 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < len(arr5); i++ {
		t.Log(arr5[i])
	}

	for index, v := range arr5 {
		t.Log(index, v)
	}
}

func TestArraySection(t *testing.T) {
	arr6 := [...]int{1, 2, 3, 4, 5, 6, 7}

	t.Log(arr6[:3])
	t.Log(arr6[:len(arr6)])

}
