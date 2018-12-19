package main

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"sync"
)

const (
	addr        = ":8080"
	width       = 500
	height      = 500
	numNeighbor = 8
)

func main() {
	rand.Seed(1234)

	sim := startSim()

	http.HandleFunc("/", homeHanlder)
	http.HandleFunc("/image.png", sim.ImageHanlder)

	log.Println("listen:", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func startSim() *Sim {
	sim := &Sim{}

	sim.syncChan = make(chan bool, width*height)

	sim.InitImages()

	// n n n
	// n p n
	// n n n

	sim.allChans = make([]chan bool, width*height)
	for i := range sim.allChans {
		sim.allChans[i] = make(chan bool, numNeighbor)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			var data PixSimData

			pixIndex := (y*width + x) * 4
			for i := 0; i < 2; i++ {
				data.pixArr[i] = sim.images[i].Pix[pixIndex : pixIndex+4]
			}

			data.neighborChans = sim.NeighborChansAt(x, y)
			data.thisChan = sim.ChanAt(x, y)

			go sim.PixSim(data)
		}
	}

	go sim.EncodeImages()

	return sim
}

type Sim struct {
	images   [2]*image.RGBA
	allChans []chan bool
	syncChan chan bool

	lock       sync.RWMutex
	encodedImg bytes.Buffer
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

type PixSimData struct {
	thisChan      chan bool
	neighborChans []chan bool
	pixArr        [2][]uint8
}

func (sim *Sim) PixSim(data PixSimData) {

	curState := data.pixArr[0][0] > 0
	index := 0

	for {
		for _, n := range data.neighborChans {
			n <- curState
		}

		numAlives := 0

		for range data.neighborChans {
			s := <-data.thisChan
			if s {
				numAlives++
			}
		}

		pix := data.pixArr[index%2]

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

		//time.Sleep(1 * time.Second)
		sim.syncChan <- curState
		index++
	}
}

func (sim *Sim) InitImages() {
	sim.images[0] = image.NewRGBA(image.Rectangle{Max: image.Point{X: width, Y: height}})
	sim.images[1] = image.NewRGBA(image.Rectangle{Max: image.Point{X: width, Y: height}})

	img := sim.images[0]
	for i := 0; i < len(img.Pix); i += 4 {
		v := uint8(rand.Intn(2) * 255)
		img.Pix[i+0] = v
		img.Pix[i+1] = v
		img.Pix[i+2] = v
		img.Pix[i+3] = 255
	}

	copy(sim.images[1].Pix, img.Pix)
}

func (sim *Sim) EncodeImages() {
	for index := 0; ; index++ {
		for i := 0; i < width*height; i++ {
			<-sim.syncChan
		}

		sim.EncodeImage(index % 2)
	}
}

func (sim *Sim) EncodeImage(imgIndex int) {
	sim.lock.Lock()
	defer sim.lock.Unlock()

	img := sim.images[imgIndex]

	sim.encodedImg.Reset()
	err := png.Encode(&sim.encodedImg, img)
	if err != nil {
		log.Fatal(err)
	}
}

var count = 0

func (sim *Sim) ImageHanlder(w http.ResponseWriter, r *http.Request) {
	count++

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)

	sim.lock.RLock()
	defer sim.lock.RUnlock()

	w.Write(sim.encodedImg.Bytes())
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
