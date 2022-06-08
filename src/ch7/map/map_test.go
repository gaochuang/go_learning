package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(len(m1))
	for k, v := range m1 {
		t.Log(k, v)
	}

	m2 := map[string]int{}
	m2["Bob"] = 23
	m2["Jane"] = 12
	for k, v := range m2 {
		t.Log(k, v)
	}

	m3 := make(map[string]int, 10)
	m3["good"] = 100
	m3["bad"] = 0
	for k, v := range m3 {
		t.Log(k, v)
	}
}

func TestNoKey(t *testing.T) {
	m1 := map[int]int{}
	//key不存在，返回0
	t.Log(m1[5])
	m1[5] = 0
	//初始化key==5的value为0，怎么判断key是否存在？
	t.Log(m1[5])

	//返回两个变量，value和一个bool类型，true代表key存在
	if v, ok := m1[3]; ok {
		t.Logf("key 3 value is %d", v)
	} else {
		t.Log("key 3 not exiting")
	}
}

func TestFactoryMode(t *testing.T) {
	m1 := map[int]func(op int) int{}
	m1[1] = func(op int) int { return op }
	m1[2] = func(op int) int { return op * op }
	m1[3] = func(op int) int { return 2 * op }

	t.Log(m1[1](2), m1[2](3), m1[3](100))
}
