package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/disintegration/imaging"
	"github.com/fatih/color"
)

func main() {
	/*ani := []rune("◴◷◶◵")
	for i := 0; i < len(ani); i++ {
		fmt.Print("\r",string(ani[i]), " ")
		time.Sleep(100*time.Millisecond)
	}*/
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
		time.Sleep(25 * time.Millisecond)
	}
	opsys := runtime.GOOS
	fmt.Println("\nOS:", opsys)

	yellsty := color.New(color.FgYellow, color.Bold)
	yellsty.Println(` 
	 ______   ______   ______   __   __       ______   ______  ______  
	/\  __ \ /\  ___\ /\  ___\ /\ \ /\ \     /\  __ \ /\  == \/\__  _\ 
	\ \  __ \\ \___  \\ \ \____\ \ \\ \ \    \ \  __ \\ \  __<\/_/\ \/ 
	 \ \_\ \_\\/\_____\\ \_____\\ \_\\ \_\    \ \_\ \_\\ \_\ \_\ \ \_\ 
	  \/_/\/_/ \/_____/ \/_____/ \/_/ \/_/     \/_/\/_/ \/_/ /_/  \/_/ 
																	   `)
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
	im, _, err := image.Decode(f)
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
	resize := []int{int(float64(size.X) * sizemul), int(float64(size.Y) * sizemul)}
	reim := imaging.Resize(im, resize[0], resize[1], imaging.Lanczos)
	reimSize := reim.Bounds().Size()

	//setup
	tone := []int{13106, 26213, 39319, 52425, 65534}
	strTone := []string{"██", "▓▓", "▒▒", "░░", "__"}
	//strTone := []string{"&&","$$","MM","OO","__"}
	strArt := ""
	fmt.Println("reading and converting pixel...")

	//lololo
	bar := []rune("[          ]")
	max := reimSize.Y - 1
	tenper := max / 10
	prog := 0
	count := 0
	//read RGBA value pixel by pixel => convert to gray value => add to string
	for y := 0; y < reimSize.Y; y++ {
		if count == tenper {
			prog++
			count = 0
		}
		percent := int(float64(float64(y)/float64(max)) * float64(100))
		switch prog {
		case 0:
			prog = 0
			fmt.Print("\rLOADING ", string(bar), percent, "%")
		default:
			bar[prog] = []rune("=")[0]
			fmt.Print("\rLOADING ", string(bar), percent, "%")
		}
		count++
		for x := 0; x < reimSize.X; x++ {
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

	//create .txt file => write string to file
	tex, err := os.Create(f.Name() + ".txt")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}
	tex.WriteString(strArt)
	//create .html file => write string to file
	html, err := os.Create(f.Name() + ".html")
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(2)
	}
	var style string
	switch opsys {
	case "windows":
		style = `style="font-family: Courier; font-size: 10px;"`
	case "linux":
		style = `style="font-family: monospace; font-size: 10px;"`
	default:
		style = ""
	}

	html.WriteString(`<button id="conv" style="font-size: 30px">Convert to image</button><pre id="ty" style="display: inline-block">  Powered by <a href="https://html2canvas.hertzen.com/">html2canvas</a></pre>
	<script>
		document.getElementById("conv").addEventListener('click', (e)=>{
		  let temp = document.createElement('pre')
		  temp.textContent = " Loading..."
		  let canvas = html2canvas(document.body)
		  canvas.then((re) => {
			let temp = document.createElement('pre')
			document.body.replaceWith(re)
		  })
		  document.getElementById("conv").replaceWith(temp)
		  document.getElementById("disp").remove()
		  document.getElementById("ty").remove()
		})
	  </script>
	  <script type="text/javascript" src="html2canvas.js"></script>
	  <pre id="disp "` + style + ">\n" + strArt + "\n</pre>")

	// ***recommend fonts: *Inversionz, courier

	fmt.Println("\nComplete...")

}
