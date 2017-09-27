package main

import (
	"bytes"
	"fmt"
	"html"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	addr := ":8080"
	http.HandleFunc("/", handler)

	log.Println("listen...", addr)
	http.ListenAndServe(addr, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	targetURL := r.URL.Query().Get("q")

	req, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	copyHeader(req.Header, r.Header)
	req.Header.Set("Accept-Encoding", "")

	//client := &http.Client{}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	copyHeader(w.Header(), resp.Header)

	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "text/html") {
		processHtml(w, resp.Body)
	} else if strings.HasPrefix(contentType, "image/") {
		processImage(w, resp.Body)
	} else {
		io.Copy(w, resp.Body)
	}
}

func processImage(w http.ResponseWriter, body io.Reader) {
	img, imgfmt, err := image.Decode(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newimg := rotate(img)

	buf := bytes.NewBuffer(nil)

	if imgfmt == "png" {
		err = png.Encode(buf, newimg)
	} else {
		err = jpeg.Encode(buf, newimg, nil)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	buf.WriteTo(w)
}

func rotate(img image.Image) image.Image {
	rect := img.Bounds()
	newimg := image.NewRGBA(rect)

	w := rect.Size().X
	h := rect.Size().Y
	for y := 0; y < h; y++ {
		y2 := rect.Min.Y + h - y - 1
		for x := 0; x < w; x++ {
			x2 := rect.Min.X + w - x - 1
			newimg.Set(x2, y2, img.At(x, y))
		}
	}

	return newimg
}

var urlRe = regexp.MustCompile(`"(http|https)://[^"]*"`)

func processHtml(w http.ResponseWriter, body io.Reader) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newData := urlRe.ReplaceAllFunc(data, func(s []byte) []byte {
		oldURL := html.UnescapeString(string(s[1 : len(s)-1]))
		newURL := "http://localhost:8080/?q=" + url.QueryEscape(oldURL)
		return ([]byte)(fmt.Sprintf(`"%s"`, newURL))
	})

	w.Header().Set("Content-Length", strconv.Itoa(len(newData)))
	w.Write(newData)
}

//type Header map[string][]string
func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
