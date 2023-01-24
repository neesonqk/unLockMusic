package main

import (
	"os"
	"time"
)

func main() {
	mpv := MPV{}
	message, code := mpv.setup()
	if code < 0 {
		panic(message)
	}

	file, err := os.OpenFile("./CantinaBand3.wav", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	mpv.playByFd(file.Fd())

	time.Sleep(4 * time.Second)
	mpv.destroy()
}
