package main

import (
	"fmt"
	"sync/atomic"
)

func MyAtomicAddInt32(addr *int32, delta int32) (newVal int32) {
	v := atomic.LoadInt32(addr)
	expected := v + newVal
	for !atomic.CompareAndSwapInt32(addr, v, expected) {
		v = atomic.LoadInt32(addr)
		expected = v + newVal
	}

	return expected
}

func MyAtomicAddZero(addr *int32, delta int32) bool {
	return MyAtomicAddInt32(addr, delta) == 0
}

var a int32

func main() {
	fmt.Println(atomic.AddInt32(&a, 100))
}
