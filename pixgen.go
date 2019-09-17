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

	"github.com/disintegration/imaging"
	"github.com/fatih/color"
)

type Config struct {
	Tone []string `json:"Tone"`
	ResizeMul float64 `json:"ResizeMul"`
}

func main() {
	//start
	fmt.Println("Starting...")
	fbar := []rune("[          ]")
	fmax := 100
	ftenper := fmax / 10
	fprog := 0
	fcount := 0
	for i := 0; i <= fmax; i++ {
		if fcount == ftenper {
			fprog++
			fcount = 0
		}
		percent := int(float64(float64(i)/float64(fmax)) * float64(100))
		switch fprog {
		case 0:
			fprog = 0
			fmt.Print("\rLOADING ", string(fbar), percent, "%")
		default:
			fbar[fprog] = []rune("=")[0]
			fmt.Print("\rLOADING ", string(fbar), percent, "%")
		}
		fcount++
		time.Sleep(4 * time.Millisecond)
	}

	yellsty := color.New(color.FgYellow)
	yellsty.Println(` 
	 ______   ______   ______   __   __       ______   ______  ______  
	/\  __ \ /\  ___\ /\  ___\ /\ \ /\ \     /\  __ \ /\  == \/\__  _\ 
	\ \  __ \\ \___  \\ \ \____\ \ \\ \ \    \ \  __ \\ \  __<\/_/\ \/ 
	 \ \_\ \_\\/\_____\\ \_____\\ \_\\ \_\    \ \_\ \_\\ \_\ \_\ \ \_\ 
	  \/_/\/_/ \/_____/ \/_____/ \/_/ \/_/     \/_/\/_/ \/_/ /_/  \/_/ 
																	   `)
	//pull config.json info
	set := Config{Tone: []string{"██", "▓▓", "▒▒", "░░", "  "}, ResizeMul: 1} //default setting
	jstr, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	} else {
		json.Unmarshal(jstr, &set) //json string to struct
	}
	fmt.Println("Preset")
	fmt.Println(" Tone:", set.Tone)
	fmt.Println(" ResizeMul:", set.ResizeMul, "\n")

	//check file in
	var finfo []os.FileInfo
	filepath.Walk("in", func(path string, info os.FileInfo, err error) error {
		finfo = append(finfo, info)
		return nil
	})
	tempin := finfo[1].Name()
	fmt.Println("Img:", tempin)
	//fmt.Println("Size(bytes):", finfo[1].Size())

	//open image
	f, err := os.Open("in/" + tempin)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}
	defer f.Close()

	//decode image
	im, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}
	//resize if needed
	size := im.Bounds().Size()
	sizemul := set.ResizeMul
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}
	resize := []int{int(float64(size.X) * sizemul), int(float64(size.Y) * sizemul)}
	reim := imaging.Resize(im, resize[0], resize[1], imaging.Lanczos)
	reimSize := reim.Bounds().Size()
	switch set.ResizeMul {
	case 1:
		fmt.Println("Size:", im.Bounds().Size().X, "x", im.Bounds().Size().Y)
	default:
		fmt.Println("Size:", im.Bounds().Size().X, "x", im.Bounds().Size().Y, "=>", reimSize.X, "x", reimSize.Y,)
	}

	//setup
	tone := []int{13106, 26213, 39319, 52425, 65534}
	strTone := set.Tone
	strArt := ""
	fmt.Println("Reading & converting pixel...")

	//laoding bar setup
	bar := []rune("[          ]")
	max := reimSize.Y - 1
	tenper := max / 10
	prog := 0
	count := 0

	//read RGBA value pixel by pixel => convert to gray value => add to string
	for y := 0; y < reimSize.Y; y++ {
		//bar compo
		if count == tenper {
			prog++
			bar[prog] = []rune("=")[0]
			count = 0
		}
		count++
		//bar compo
		for x := 0; x < reimSize.X; x++ {
			//bar compo
			percent := int(float64(float64(y)/float64(max)) * float64(100))
			fmt.Print("\rLOADING ", string(bar), percent, "% [y: ", y+1, "/", reimSize.Y, "] [x: ", x+1, "/", reimSize.X, "]")
			/*switch y {
			case reimSize.Y-1:
				fmt.Print("\rLOADING ", string(bar), percent, "%                                                               ")
			default:
				fmt.Print("\rLOADING ", string(bar), percent, "% [y: ", y+1, "/", reimSize.Y,"] [x: ", x+1, "/", reimSize.X, "]")
			}*/
			//bar compo
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

	//create text.txt file => write string to file
	tex, err := os.Create("out/text.txt")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}
	tex.WriteString(strArt)

	//create text.html file => write string to file
	html, err := os.Create("out/text.html")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}
	html.WriteString(`<pre id="disp" style="font-family: Courier; font-size: 10px;"` + ">\n" + strArt + "\n</pre>\n")

	//create img.html file => write string to file
	htmlimg, err := os.Create("out/img.html")
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

	// ***recommend fonts: *Inversionz, courier

	fmt.Println("\nComplete...")
}
