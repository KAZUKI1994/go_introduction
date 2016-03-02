package main

import "image"
import "image/gif"
import "os"

func main(){
	files := []string{"g1.gif", "g2.gif", "g3.gif", "g2.gif"}

	// 各フレームの画像をGIFで読み込んでoutGifを構築する
	outGif := &gif.GIF{}
	for _, name := range files{
		f, _ := os.Open(name)
		inGif, _ := gif.Decode(f)
		f.Close()

		outGif.Image = append(outGif.Image, inGif.(*image.Paletted))
		outGif.Delay = append(outGif.Delay, 0)
	}

	// out.gifに保存する
	f, _ := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.EncodeAll(f, outGif)
}
