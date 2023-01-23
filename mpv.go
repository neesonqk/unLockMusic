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
char* longToChar(unsigned long source) {
	return (char*)source;
}
*/
import "C"

import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

type MPV struct {
	instance *C.mpv_handle
	mutex    sync.Mutex
	running  bool
}

func (mpv *MPV) setStringOption(key string, value string) (string, int) {
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))

	return setOption(mpv.instance, key, unsafe.Pointer(&cValue), C.MPV_FORMAT_STRING)
}

func (mpv *MPV) setIntOption(key string, value int) (string, int) {
	cValue := C.int64_t(value)

	return setOption(mpv.instance, key, unsafe.Pointer(&cValue), C.MPV_FORMAT_INT64)
}

func (mpv *MPV) setBoolOption(key string, value bool) (string, int) {
	cValue := C.int(0)
	if value {
		cValue = 1
	}

	return setOption(mpv.instance, key, unsafe.Pointer(&cValue), C.MPV_FORMAT_FLAG)
}

func setOption(instance *C.mpv_handle, key string, value unsafe.Pointer, format C.mpv_format) (string, int) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	status := C.mpv_set_option(instance, cKey, format, value)
	return C.GoString(C.mpv_error_string(status)), int(status)
}

func (mpv *MPV) create() {
	mpv.instance = C.mpv_create()
}

func (mpv *MPV) run() {

	// cKey := C.CString("softvol")
	// cValue := C.CString("yes")
	// defer C.free(unsafe.Pointer(cKey))
	// C.mpv_set_option(mpv, cKey, C.MPV_FORMAT_STRING, unsafe.Pointer(&cValue))

	C.mpv_initialize(mpv.instance)

	major := C.mpv_client_api_version()
	fmt.Println(major)

	var cmd [3]string
	cmd[0] = "loadfile"
	cmd[1] = "file:///Users/neeson/Downloads/SoundHelix-Song-13.mp3"
	cmd[2] = "replace"
	// cmd[3] = "start=+100,vid=no"
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

	// cStr4 := C.CString(cmd[3])
	// C.setArrayString(cArray, C.int(3), cStr4)
	// defer C.free(unsafe.Pointer(cStr4))

	var e = C.mpv_command(mpv.instance, cArray)
	fmt.Println(int(e))
	fmt.Println(C.GoString(C.mpv_error_string(e)))
	// mpv_terminate_destroy(mpv)

	time.Sleep(20 * time.Second)
	C.mpv_terminate_destroy(mpv.instance)
}

func (mpv *MPV) testFd(fd uintptr) {
	C.mpv_initialize(mpv.instance)

	major := C.mpv_client_api_version()
	fmt.Println(major)

	fmt.Println(fmt.Sprint(fd))

	var cmd [3]string
	cmd[0] = "loadfile"
	cmd[1] = "fd://" + fmt.Sprint(fd)
	cmd[2] = "replace"
	// cmd[3] = "start=+100,vid=no"
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

	// cStr4 := C.CString(cmd[3])
	// C.setArrayString(cArray, C.int(3), cStr4)
	// defer C.free(unsafe.Pointer(cStr4))

	var e = C.mpv_command(mpv.instance, cArray)
	fmt.Println(int(e))
	fmt.Println(C.GoString(C.mpv_error_string(e)))
	// mpv_terminate_destroy(mpv)

	time.Sleep(20 * time.Second)
	C.mpv_terminate_destroy(mpv.instance)
}
