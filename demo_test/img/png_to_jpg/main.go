package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

func main() {
	originalFilePath := "/Users/shikuanxu/Documents/父母申请PR/WX20230212-201512.png"
	createFilePath := "/Users/shikuanxu/Documents/父母申请PR/WX20230212-201512.jpeg"
	pngImgFile, err := os.Open(originalFilePath)

	if err != nil {
		fmt.Println("PNG-file.png file not found!")
		os.Exit(1)
	}

	defer pngImgFile.Close()

	// create image from PNG file
	imgSrc, err := png.Decode(pngImgFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create a new Image with the same dimension of PNG image
	newImg := image.NewRGBA(imgSrc.Bounds())

	// we will use white background to replace PNG's transparent background
	// you can change it to whichever color you want with
	// a new color.RGBA{} and use image.NewUniform(color.RGBA{<fill in color>}) function

	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// paste PNG image OVER to newImage
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)

	// create new out JPEG file
	jpgImgFile, err := os.Create(createFilePath)

	if err != nil {
		fmt.Println("Cannot create JPEG-file.jpg !")
		fmt.Println(err)
		os.Exit(1)
	}

	defer jpgImgFile.Close()

	var opt jpeg.Options
	opt.Quality = 80

	// convert newImage to JPEG encoded byte and save to jpgImgFile
	// with quality = 80
	err = jpeg.Encode(jpgImgFile, newImg, &opt)

	//err = jpeg.Encode(jpgImgFile, newImg, nil) -- use nil if ignore quality options

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Converted PNG file to JPEG file")
}
