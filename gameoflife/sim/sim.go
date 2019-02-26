package sim

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/rand"
	"sync"
)

const numNeighbor = 8

func New(width, height int) *Sim {
	sim := &Sim{
		width:  width,
		height: height,
	}

	sim.frameCh[0] = make(chan struct{})

	sim.InitImages()

	// n n n
	// n p n
	// n n n

	sim.allChans = make([]chan Msg, width*height)
	for i := range sim.allChans {
		sim.allChans[i] = make(chan Msg, numNeighbor)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			var data PixSimData

			pixIndex := (y*width + x) * 4
			data.pixArr = sim.images.Pix[pixIndex : pixIndex+4]

			data.neighborChans = sim.NeighborChansAt(x, y)
			data.thisChan = sim.ChanAt(x, y)

			go sim.PixSim(data)
		}
	}

	sim.inputCh = make(chan func(), 100)

	return sim
}

type Sim struct {
	width, height int

	images   *image.RGBA
	allChans []chan Msg

	lock       sync.RWMutex
	encodedImg bytes.Buffer

	//syncChan chan bool
	curFrame int
	frameWg  sync.WaitGroup
	frameCh  [2]chan struct{}

	inputCh chan func()
}

type Msg struct {
	state bool
	frame int
}

func (sim *Sim) ChanAt(x, y int) chan Msg {
	return sim.allChans[y*sim.width+x]
}

func (sim *Sim) NeighborChansAt(x, y int) []chan Msg {
	chans := make([]chan Msg, 0, numNeighbor)

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}

			x2 := (x + dx + sim.width) % sim.width
			y2 := (y + dy + sim.height) % sim.height

			chans = append(chans, sim.ChanAt(x2, y2))
		}
	}

	return chans
}

type PixSimData struct {
	thisChan      chan Msg
	neighborChans []chan Msg
	pixArr        []uint8
}

func (sim *Sim) PixSim(data PixSimData) {

	for frame := 0; ; frame++ {
		<-sim.frameCh[frame%2]

		curState := data.pixArr[0] > 0

		for _, n := range data.neighborChans {
			n <- Msg{state: curState, frame: frame}
		}

		numAlives := 0

		for range data.neighborChans {
			s := <-data.thisChan
			if s.frame != frame {
				panic(fmt.Sprintf("incorrect input frame: input=%v, cur=%v", s.frame, frame))
			}

			if s.state {
				numAlives++
			}
		}

		pix := data.pixArr

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

		/*

		1. main: wg.Add(w*h)
		2. main: (frame+1) make(chan..)

		3. block -> unblock - (cur frame)chan.recv/close
		   cell: recv
		   main: close

		4. cell: Done x w*h
		5. main: wg.Wait - waitgroup (w*h)

		 */

		//time.Sleep(1 * time.Second)

		sim.frameWg.Done()
	}
}

func (sim *Sim) InitImages() {
	sim.images = image.NewRGBA(image.Rectangle{Max: image.Point{X: sim.width, Y: sim.height}})

	img := sim.images
	for i := 0; i < len(img.Pix); i += 4 {
		v := uint8(rand.Intn(2) * 255)
		img.Pix[i+0] = v
		img.Pix[i+1] = v
		img.Pix[i+2] = v
		img.Pix[i+3] = 255
	}

}

func (sim *Sim) EncodeImages() {
	for index := 0; ; index++ {
		sim.WaitForFrame()

		sim.ProcessInputs()

		sim.EncodeImage()
	}
}

func (sim *Sim) WaitForFrame() {

	sim.frameWg.Add(sim.width * sim.height)

	sim.frameCh[(sim.curFrame+1)%2] = make(chan struct{})

	close(sim.frameCh[sim.curFrame%2])

	sim.curFrame++

	sim.frameWg.Wait()
}

func (sim *Sim) EncodeImage() {
	sim.lock.Lock()
	defer sim.lock.Unlock()

	img := sim.images

	sim.encodedImg.Reset()
	err := png.Encode(&sim.encodedImg, img)
	if err != nil {
		log.Fatal(err)
	}
}

var count = 0

func (sim *Sim) WriteImage(w io.Writer) error {
	count++

	sim.lock.RLock()
	defer sim.lock.RUnlock()

	_, err := w.Write(sim.encodedImg.Bytes())
	return err
}

func (sim *Sim) SetPixel(x, y int) {
	w := sim.width
	x = ((x % w) + w) % w

	h := sim.height
	y = ((y % h) + h) % h

	sim.images.Set(x, y, color.White)
}

func (sim *Sim) OnClick(x, y int) error {

	go func() {
		sim.inputCh <- func() {
			sim.SetPixel(x-1, y)
			sim.SetPixel(x, y)
			sim.SetPixel(x+1, y)
		}
	}()

	return nil
}

func (sim *Sim) ProcessInputs() {
	for {
		select {
		case f := <-sim.inputCh:
			f()
		default:
			return
		}
	}
}
