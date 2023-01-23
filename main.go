package main

import (
	"fmt"
	"os"
)

func main() {
	mpv := MPV{}

	mpv.create()
	// s := mpv.setStringOption("softvol", "yes")
	// fmt.Println(s)

	ss, code := mpv.setBoolOption("resume-playback", false)
	fmt.Println("setFlag resume-playback", ss, code)

	ss, code = mpv.setBoolOption("cache", true)
	fmt.Println("setFlag cache", ss, code)

	ss, code = mpv.setIntOption("cache-secs", 160) // 10 seconds
	fmt.Println("setInt cache-default", ss, code)

	// mpv.run()

	file, err := os.OpenFile("/Users/neeson/Downloads/SoundHelix-Song-13.mp3", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fd := file.Fd()
	mpv.testFd(fd)
}
