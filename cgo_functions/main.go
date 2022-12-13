package main

import (
	"fmt"
	"sync"
)

/*
extern void go_call_back(int foo, int p);

static inline void CallMyFunc(int foo) {
	go_call_back(foo, 5);
}
*/
import "C"

//export go_call_back
func go_call_back(foo C.int, p C.int) {
	fn := lookup(int(foo))
	fn(p)
}

func MyCallback(x C.int) {
	fmt.Println("callback with", x)
}

func Example() {
	i := register(MyCallback)
	C.CallMyFunc(C.int(i))
	unregister(i)
}

func main() {
	Example()
}

// /////////////////////////
var mu sync.Mutex
var index int
var fns = make(map[int]func(C.int))

func register(fn func(C.int)) int {
	mu.Lock()
	defer mu.Unlock()
	index++
	for fns[index] != nil {
		index++
	}
	fns[index] = fn
	return index
}

func lookup(i int) func(C.int) {
	mu.Lock()
	defer mu.Unlock()
	return fns[i]
}

func unregister(i int) {
	mu.Lock()
	defer mu.Unlock()
	delete(fns, i)
}
