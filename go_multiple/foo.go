package main

/*
#include <stdio.h>
#include <stdlib.h>

void sprintFoo(char* c) {
	printf("foo: %s\n", c);
}
*/
import "C"

import "unsafe"

func PrintFoo() {
	foo := C.CString("Hello ! I'm Foo")
	C.sprintFoo(foo)
	// release foo variable
	C.free(unsafe.Pointer(foo))
}
