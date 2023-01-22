package main

/*
#cgo CFLAGS: -I./lib
#cgo LDFLAGS: -L./lib -lmpv.1 -Wl,-rpath,./lib
#include <stdio.h>
#include <stdlib.h>
#include "client.h"

//** Helper functions **
char** makeCharArray(int size) {
	return calloc(sizeof(char*), size);
}
void setArrayString(char** a, int i, char* s) {
	a[i] = s;
}
*/
import "C"

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("Hello")
	mpv := C.mpv_create()
	fmt.Println(mpv)

	C.mpv_initialize(mpv)
	// mpv.setOptionInt("volume", C.int(5))

	var cmd [4]string
	cmd[0] = "loadfile"
	cmd[1] = "https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3"
	cmd[2] = "replace"
	cmd[3] = "start=+100,vid=no"
	// cmd[2] = ""

	cArray := C.makeCharArray(C.int(len(cmd) + 1))
	defer C.free(unsafe.Pointer(cArray))

	cStr := C.CString(cmd[0])
	C.setArrayString(cArray, C.int(0), cStr)
	defer C.free(unsafe.Pointer(cStr))

	cStr2 := C.CString(cmd[1])
	C.setArrayString(cArray, C.int(1), cStr2)
	defer C.free(unsafe.Pointer(cStr2))

	cStr3 := C.CString(cmd[2])
	C.setArrayString(cArray, C.int(2), cStr3)
	defer C.free(unsafe.Pointer(cStr3))

	cStr4 := C.CString(cmd[3])
	C.setArrayString(cArray, C.int(3), cStr4)
	defer C.free(unsafe.Pointer(cStr4))

	var e = C.mpv_command(mpv, cArray)
	fmt.Println(int(e))
	fmt.Println(C.GoString(C.mpv_error_string(e)))
	// mpv_terminate_destroy(mpv)

	time.Sleep(20 * time.Second)
	C.mpv_terminate_destroy(mpv)
}
