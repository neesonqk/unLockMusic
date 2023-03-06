package player

import (
	"zhiyin/boya/v2/player/mpv_lib"

	"github.com/ying32/govcl/vcl"
)

// ::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {

}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {

	mpv := mpv_lib.GetInstance()

	// mpv := mpv_lib.MPV{}

	// message, code := mpv.Setup()
	// if code < 0 {
	// 	panic(message)
	// }

	// file, err := os.OpenFile("./CantinaBand3.wav", os.O_RDONLY, 0644)
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	mpv.PlayByPath("./CantinaBand3.wav")

	// // time.Sleep(14 * time.Second)
	// mpv.Destroy()

	// fmt.Println("finished")

}
