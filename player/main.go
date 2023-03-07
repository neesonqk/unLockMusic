package main

import (
	"fmt"
	"zhiyin/boya/v2/player/ui/player"

	"github.com/ying32/govcl/vcl"
)

// #include "lib.h"
// #include <stdlib.h>
//#include <locale.h>
//void setLocal() { setlocale(LC_NUMERIC, "C"); }
import "C"

func main() {
	x := C.Add(1, 2)
	fmt.Println(x)

	vcl.Application.SetScaled(true)
	vcl.Application.SetTitle("Boya Music Player")
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	// vcl.Application.CreateForm(&form.Form1)
	vcl.Application.CreateForm(&player.Form1)

	// https://github.com/mpv-player/mpv-examples/blob/bd0b42e5bd47c22592760a244f80e49ec0222892/libmpv/qt/qtexample.cpp?rgh-link-date=2019-10-29T20%3A56%3A11Z#L217
	// check this place, when vcl get imported, value from setLocale is overridden,
	// so set it back again, as this is required by mpv library.
	C.setLocal()

	vcl.Application.Run()
}
