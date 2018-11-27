package main

import (
	"io/ioutil"
	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
)

func main() {
	var file, _ = ioutil.ReadFile("杨坤 - 那一天.mp3")
	dec, data, _ := minimp3.DecodeFull(file)
	player, _ := oto.NewPlayer(dec.SampleRate, dec.Channels, 2, 1024*8)
	player.Write(data)
	}
