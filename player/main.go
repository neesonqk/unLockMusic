package main

import (
	"zhiyin/boya/v2/player/ui/form"

	_ "github.com/ying32/govcl/pkgs/macapp"
	"github.com/ying32/govcl/vcl"
)

func main() {

	vcl.Application.SetScaled(true)
	vcl.Application.SetTitle("form")
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&form.Form1)

	// mpv := MPV{}

	// message, code := mpv.setup()
	// if code < 0 {
	// 	panic(message)
	// }

	// file, err := os.OpenFile("./CantinaBand3.wav", os.O_RDONLY, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// mpv.playByFd(file.Fd())

	// time.Sleep(4 * time.Second)
	// mpv.destroy()

	vcl.Application.Run()
}
