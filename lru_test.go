//----------------
//Func  :LRU Container Testing
//Author: xjh
//Date  : 2018/09/10
//Note  :
//----------------
package lru

import (
	"fmt"
	"testing"
)

func TestLruContainer_PushFront(t *testing.T) {
	lc := NewLruContainer(10)
	for idx := 0; idx < 100; idx++ {
		lc.PushFront(idx%10, idx)
	}
	fmt.Println(lc.Len())

	fmt.Println(lc.Get(5))
	fmt.Println(lc.Get(4))

	lc.Del(7)

	for idx := 0; idx < 10; idx++ {
		fmt.Println(lc.PopBack())
	}

	fmt.Println(lc.Get(4))
	fmt.Println(lc.Get(5))

	fmt.Println(lc.Len())
}

var glc = NewLruContainer(1000)
var idx int

func BenchmarkLruContainer_PushFront(b *testing.B) {
	glc.PushFront(idx, idx)
}

func BenchmarkLruContainer_Get(b *testing.B) {
	glc.Get(idx)
}
