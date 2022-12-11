package main

/*
#include <stdio.h>
#include <stdlib.h>
void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"
import "unsafe"

func main() {
	Example()
}

func Example() {
	cs := C.CString("Hello")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}
