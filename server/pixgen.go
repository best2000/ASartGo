package main

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/fatih/color"
)

type Config struct {
	Tone []string `json:"Tone"`
	ResizeMul string `json:"ResizeMul"`
}

func elapsed() func() {
    start := time.Now()
    return func() {
        fmt.Println("\nTimer: ", time.Since(start))
    }
}

func main() {
	color.Set(color.FgHiGreen)
	defer color.Unset()
	defer elapsed()()
	
	//pull config.json info
	set := Config{Tone: []string{"██", "▓▓", "▒▒", "░░", "  "}, ResizeMul: "1"} //default setting
	jstr, err := ioutil.ReadFile("conin/config.json")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	} else {
		json.Unmarshal(jstr, &set) //json string to struct
	}
	os.Remove("conin/config.json") //remove readed file
	fmt.Println("Preset")
	fmt.Println(" Tone:", set.Tone)
	fmt.Println(" ResizeMul:", set.ResizeMul, "\n")
	

	//check file in
	var finfo []os.FileInfo
	filepath.Walk("in", func(path string, info os.FileInfo, err error) error {
		finfo = append(finfo, info)
		return nil
	})
	fname := finfo[1].Name()
	fmt.Println("Img:", fname)

	//open image
	f, err := os.Open("in/" + fname)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}

	//decode image
	im, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}

	f.Close()
	os.Remove("in/"+fname) //remove readed file

	//resize if needed
	size := im.Bounds().Size()
	sizemul, _ := strconv.ParseFloat(set.ResizeMul, 64)
	sizemul = sizemul/100
	resize := []int{int(float64(size.X) * sizemul), int(float64(size.Y) * sizemul)}
	reim := imaging.Resize(im, resize[0], resize[1], imaging.Lanczos)
	reimSize := reim.Bounds().Size()
	fmt.Println("Size:", im.Bounds().Size().X, "x", im.Bounds().Size().Y, "=>", reimSize.X, "x", reimSize.Y,)
	

	//setup
	tone := []int{13106, 26213, 39319, 52425, 65534}
	strTone := set.Tone
	strArt := ""
	fmt.Println("Reading & converting pixel...")

	//prgess show setup
	maxpix := float64(reimSize.X * reimSize.Y)
	var pxcount float64 
	var percent int
	
	//read RGBA value pixel by pixel => convert to gray value => add to string
	for y := 0; y < reimSize.Y; y++ {
		for x := 0; x < reimSize.X; x++ {
			//progess compo
			pxcount++
			percent = int(pxcount/maxpix*100)
			fmt.Print("\rLoading: [", percent, "%] [px: ", pxcount, "/", maxpix , "] [y: ", y+1, "/", reimSize.Y, "] [x: ", x+1, "/", reimSize.X, "]")
			//progess compo end
			r, g, b, _ := reim.At(x, y).RGBA()
			gray := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
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

	/*//create text.html file => write string to file
	html, err := os.Create("out/text.html")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}
	html.WriteString(`<pre id="disp" style="font-family: Courier; font-size: 10px;"` + ">\n" + strArt + "\n</pre>\n")*/

	//create img.html file => write string to file
	htmlimg, err := os.Create("out/"+fname+".html")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}
	htmlimg.WriteString(
		`<pre id="disp" style="font-family: Courier; font-size: 10px;"` + ">\n" + strArt + "\n</pre>\n")

	//html2canvas
	jsbyte, _ := ioutil.ReadFile("html2canvas.min.js")
	htmlimg.WriteString("<script>\n" + string(jsbyte) + "\n</script>\n")

	htmlimg.WriteString(`
	<script>	  
		let canvas = html2canvas(document.body)
  		canvas.then((re) => {
		document.body.replaceWith(re)
  		})
	</script>`)
	fmt.Println("\nImg out:", fname+".html")

	// ***recommend fonts: *Inversionz, courier
}
