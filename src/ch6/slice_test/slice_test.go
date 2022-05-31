package slicetest

import "testing"

func TestSlice1(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)

	t.Log(len(s0), cap(s0))
}

func TestSlice2(t *testing.T) {

	s2 := make([]int, 3, 5)
	//t.Log(s2[4]) 报错
	t.Log(len(s2), cap(s2))
	t.Log(s2)
	s2 = append(s2, 888)
	t.Log(s2)
}

func TestSlice3(t *testing.T) {
	s3 := []int{}
	for i := 0; i < 10; i++ {
		s3 = append(s3, i)
		t.Log(len(s3), cap(s3))
	}
}

func TestSlice4(t *testing.T) {
	s4 := make([]int, 2, 2)
	for i := 0; i < 10; i++ {
		s4 = append(s4, i)
		t.Log(len(s4), cap(s4))
	}
}
