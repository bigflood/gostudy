package main

import (
	"image"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	addr        = ":8080"
	width       = 500
	height      = 500
	numNeighbor = 8
)

var (
	img = image.NewRGBA(image.Rectangle{Max: image.Point{X: width, Y: height}})
)

func initImg() {
	for i := 0; i < len(img.Pix); i += 4 {
		v := uint8(rand.Intn(2) * 255)
		img.Pix[i+0] = v
		img.Pix[i+1] = v
		img.Pix[i+2] = v
		img.Pix[i+3] = 255
	}
}

func startSim() {
	initImg()

	sim := &Sim{}

	// n n n
	// n p n
	// n n n

	sim.allChans = make([]chan bool, width*height)
	for i := range sim.allChans {
		sim.allChans[i] = make(chan bool, numNeighbor)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pixIndex := (y*width + x) * 4
			pix := img.Pix[pixIndex : pixIndex+4]

			neighborChans := sim.NeighborChansAt(x, y)
			thisChan := sim.ChanAt(x, y)

			go sim.PixSim(
				thisChan,
				neighborChans,
				pix,
			)

		}
	}
}

type Sim struct {
	allChans []chan bool
}

func (sim *Sim) ChanAt(x, y int) chan bool {
	return sim.allChans[y*width+x]
}

func (sim *Sim) NeighborChansAt(x, y int) []chan bool {
	chans := make([]chan bool, 0, numNeighbor)

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			x2 := (x + dx + width) % width
			y2 := (y + dy + height) % height

			chans = append(chans, sim.ChanAt(x2, y2))
		}
	}

	return chans
}

func (sim *Sim) PixSim(thisChan chan bool, neighborChans []chan bool, pix []uint8) {
	curState := pix[0] > 0

	for {
		for _, n := range neighborChans {
			n <- curState
		}

		numAlives := 0

		for range neighborChans {
			s := <-thisChan
			if s {
				numAlives++
			}
		}

		switch numAlives {
		case 2:
		case 3:
			curState = true
			pix[0] = 255
			pix[1] = 255
			pix[2] = 255
		default:
			curState = false
			pix[0] = 0
			pix[1] = 0
			pix[2] = 0
		}

		time.Sleep(1*time.Second)
	}
}

func main() {
	startSim()

	http.HandleFunc("/", homeHanlder)
	http.HandleFunc("/image.png", imageHanlder)

	log.Println("listen:", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

var count = 0

func imageHanlder(w http.ResponseWriter, r *http.Request) {
	count++

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	png.Encode(w, img)
}

func homeHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(homeHtml))
}

const homeHtml = `<html>
<head>
</head>
<body>
<img id="img" src="/image.png">

<script>
window.onload = function() {
    var image = document.getElementById("img");

    function updateImage() {
        image.src = image.src.split("?")[0] + "?" + new Date().getTime();
    }

    setInterval(updateImage, 1000);
}
</script>

</body>
</html>
`
