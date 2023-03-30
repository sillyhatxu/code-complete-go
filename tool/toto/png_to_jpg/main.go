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

func test(input, output string) {
	//pngImgFile, err := os.Open("./PNG-file.png")
	pngImgFile, err := os.Open(input)
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
	//jpgImgFile, err := os.Create("./JPEG-file.jpg")
	jpgImgFile, err := os.Create(output)
	if err != nil {
		fmt.Println("Cannot create JPEG-file.jpg !")
		fmt.Println(err)
		os.Exit(1)
	}
	defer jpgImgFile.Close()
	var opt jpeg.Options
	opt.Quality = 99
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

func main() {
	path := "/Users/shikuanxu/Documents/父母申请PR/LTVP/approve/untitled folder/"
	test(path+"1-1.png", path+"1-1.jpg")
	test(path+"1-2.png", path+"1-2.jpg")
	test(path+"1-3.png", path+"1-3.jpg")
	test(path+"2-1.png", path+"2-1.jpg")
	test(path+"2-2.png", path+"2-2.jpg")
	//inputs := []string{
	//	"/Users/xushikuan/Documents/doc/Citizen/6. 雇主信-工资单-收入-雇用时间/1. February 2021.png",
	//	"/Users/xushikuan/Documents/doc/Citizen/6. 雇主信-工资单-收入-雇用时间/2. January 2021.png",
	//	"/Users/xushikuan/Documents/doc/Citizen/6. 雇主信-工资单-收入-雇用时间/3. December 2020.png",
	//	"/Users/xushikuan/Documents/doc/Citizen/6. 雇主信-工资单-收入-雇用时间/4. November 2020.png",
	//	"/Users/xushikuan/Documents/doc/Citizen/6. 雇主信-工资单-收入-雇用时间/5. October 2020.png",
	//	"/Users/xushikuan/Documents/doc/Citizen/6. 雇主信-工资单-收入-雇用时间/6. September 2020.png",
	//}
	//outputs := []string{
	//	"/Users/xushikuan/Documents/doc/Citizen/6. 雇主信-工资单-收入-雇用时间/1. February 2021.jpg",
	//	"/Users/xushikuan/Documents/doc/Citizen/6. 雇主信-工资单-收入-雇用时间/2. January 2021.jpg",
	//	"/Users/xushikuan/Documents/doc/Citizen/6. 雇主信-工资单-收入-雇用时间/3. December 2020.jpg",
	//	"/Users/xushikuan/Documents/doc/Citizen/6. 雇主信-工资单-收入-雇用时间/4. November 2020.jpg",
	//	"/Users/xushikuan/Documents/doc/Citizen/6. 雇主信-工资单-收入-雇用时间/5. October 2020.jpg",
	//	"/Users/xushikuan/Documents/doc/Citizen/6. 雇主信-工资单-收入-雇用时间/6. September 2020.jpg",
	//}
	//for i, v := range inputs {
	//	test(v, outputs[i])
	//}
}
