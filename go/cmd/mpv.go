package main

/*
#cgo CFLAGS: -I./includes
#cgo LDFLAGS: -L./includes -lmpv.1 -Wl,-rpath,./includes
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

func (mpv *MPV) setup() (string, int) {
	mpv.instance = C.mpv_create()

	mpv.setBoolOption("resume-playback", false)
	mpv.setBoolOption("cache", true)
	mpv.setIntOption("cache-secs", 160) // 10 seconds
	mpv.setBoolOption("video", false)

	status := C.mpv_initialize(mpv.instance)
	return mpvResult(status)
}

func (mpv *MPV) destroy() {
	C.mpv_terminate_destroy(mpv.instance)
}

func (mpv *MPV) createPlayList() {

}

func (mpv *MPV) playByFd(fd uintptr) (string, int) {
	if mpv.instance == nil {
		return "MPV is not setup!", -1
	}
	cmd := []string{"loadfile", "fd://" + fmt.Sprint(fd), "replace"}
	status := execCommand(mpv.instance, cmd)
	return mpvResult(status)
}

func (mpv *MPV) playByPath(filepath string) (string, int) {
	if mpv.instance == nil {
		return "MPV is not setup!", -1
	}

	cmd := []string{"loadfile", "file://" + filepath, "replace"}
	status := execCommand(mpv.instance, cmd)
	return mpvResult(status)
}

func setOption(instance *C.mpv_handle, key string, value unsafe.Pointer, format C.mpv_format) (string, int) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	status := C.mpv_set_option(instance, cKey, format, value)
	return mpvResult(status)
}

func execCommand(instance *C.mpv_handle, command []string) C.int {
	cArray := C.makeCharArray(C.int(len(command)))
	for i, s := range command {
		C.setArrayString(cArray, C.int(i), C.CString(s))
	}

	return C.mpv_command(instance, cArray)
}

func mpvResult(status C.int) (string, int) {
	return C.GoString(C.mpv_error_string(status)), int(status)
}
