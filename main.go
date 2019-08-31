package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"sort"
	"bytes"
)


func main() {
	f, err := os.Open("bg.jpg")
	fmt.Println("err:", err)
	defer f.Close()

	im, typ, err := image.Decode(f)
	fmt.Println("err:", err)

	size := im.Bounds().Size()
	rect := image.Rect(0,0,size.X,size.Y)
	imgay := image.NewGray16(rect)

	a := []int{13106, 26213, 39319, 52425, 65534}
	shade := []string{"#","D","O","/","_"}
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			r, g, b, _ := im.At(x,y).RGBA()
			gray := 0.299 * float64(r) +  0.587 * float64(g) + 0.114 * float64(b)
			grayInt := int(gray)
			
			x := grayInt
			i := sort.Search(len(a), func(i int) bool { return a[i] <= x })
			//write buffer to keep text then write to text file !!!!!!!!!!!

			c := color.Gray16{Y: uint16(grayInt)}
			imgay.Set(x, y, c) 
		}
	}
	///d/sad/sadsad/sad/sad/
	tex, err := os.Create("ASart.txt")
	fmt.Println("err:", err)
	tex.WriteString(buff.string())

	//create/save grayscale to new image
	nim, err := os.Create("imgay2.png")
	fmt.Println("err:", err)
	defer nim.Close()
	
	switch typ {
	case "jpeg":
		err = jpeg.Encode(nim, imgay, nil)
		fmt.Println("err:", err)
	case "png":
		err = png.Encode(nim, imgay)
		fmt.Println("err:", err)
	default:
		fmt.Println("err: No TYPE!")
	}
	
	
}	
