package main

import (
	"fmt"
	"sync/atomic"
)

func atomicChanges() {
	// Antes da versao 1.19
	//var i int64
	//atomic.AddInt64(&i, 10)
	//fmt.Println(atomic.LoadInt64(&i))

	//Depois da versao 1.19
	var i atomic.Int64
	i.Add(10)
	fmt.Println(i.Load())
}
