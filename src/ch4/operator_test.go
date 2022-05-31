package ch4_test

import "testing"

func TestArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5, 6}
	b := [...]int{6, 7, 8, 9, 10, 11}
	c := [...]int{1, 2, 3, 4, 5, 6}

	t.Log(a == b)
	t.Log(a == c)
}
