/*
goroutine并发模式--只运行一次

*/
package ch23_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var sigleInstalce *Singleton

var once sync.Once

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create obj")
		sigleInstalce = new(Singleton)
	})
	return sigleInstalce
}

//单例 懒汉模式 线程安全
func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%+v\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
