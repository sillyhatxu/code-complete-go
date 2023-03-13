package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/png"
	"log"
	"os"
)

func main() {
	file_path := "/Users/shikuanxu/Documents/父母申请PR/我/徐毕业证学位证公证书拍照/学位公证1.jpg"

	fmt.Println("Hello World")
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(800, 0, img, resize.NearestNeighbor)

	out, err := os.Create("CHINESE_Renunciation_Certificate_2.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}
