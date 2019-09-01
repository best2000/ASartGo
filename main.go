package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"github.com/disintegration/imaging"
	"strconv"
)


func main() {

	fmt.Print("Image path: ")
    var path string
	fmt.Scanln(&path)
	
	//open image
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}
	defer f.Close()

	//decode image
	im, typ, err := image.Decode(f)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}

	//resize if needed
	fmt.Print("Resize mul: ")
    var sizemulstr string
	fmt.Scanln(&sizemulstr) 
	size := im.Bounds().Size()
	sizemul, err := strconv.ParseFloat(sizemulstr, 64)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}
	resize := []int{int(float64(size.X)*sizemul), int(float64(size.Y)*sizemul)}
	reim := imaging.Resize(im, resize[0], resize[1], imaging.Lanczos)
	reimSize := reim.Bounds().Size()

	//setup
	tone := []int{13106, 26213, 39319, 52425, 65534}
	strTone := []string{"#","o","-","O"," "}
	strArt := ""

	//read RGBA value pixel by pixel => convert to gray value => add to string 
	for y := 0; y < reimSize.Y; y++ {
		for x := 0; x < reimSize.X; x++ {
			r, g, b, _ := reim.At(x,y).RGBA()
			gray := 0.299 * float64(r) +  0.587 * float64(g) + 0.114 * float64(b)
			grayInt := int(gray)
			
			if grayInt <= tone[0] {
				strArt += strTone[0]
			} else if grayInt > tone[0] && grayInt <= tone[1] {
				strArt += strTone[1]
			} else if grayInt > tone[1] && grayInt <= tone[2] {
				strArt += strTone[2]
			} else if grayInt > tone[2] && grayInt <= tone[3] {
				strArt += strTone[3]
			} else {
				strArt += strTone[4]
			}
		}
		strArt += "\n"
	}

	//create .txt file => write string to file
	tex, err := os.Create(f.Name() + ".txt")
	if err!= nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}
	tex.WriteString(strArt)

	// ***recommend fonts: *Inversionz, all monospace 

	//create/save resize to new image
	nim, err := os.Create("re" + f.Name())
	fmt.Println("err:", err)
	defer nim.Close()
	
	switch typ {
	case "jpeg":
		err = jpeg.Encode(nim, reim, nil)
		fmt.Println("err:", err)
	case "png":
		err = png.Encode(nim, reim)
		fmt.Println("err:", err)
	default:
		fmt.Println("err: No TYPE!")
	}
	
	
}	
