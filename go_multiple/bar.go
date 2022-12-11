package main

/*
#include <stdio.h>
#include <stdlib.h>

void sprintBar(char* c) {
	printf("bar: %s\n", c);
}
*/
import "C"
import "unsafe"

func PrintBar() {
	bar := C.CString("Hello ! I'm Bar")
	C.sprintBar(bar)
	// release foo variable
	C.free(unsafe.Pointer(bar))
}
