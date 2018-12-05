package main

import (
	"image"
	"image/png"
	"log"
	"net/http"
)

const (
	addr   = ":8080"
	width  = 500
	height = 500
)

func main() {
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

	img := image.NewRGBA(image.Rectangle{Max: image.Point{X:width, Y:height}})
	for i := 0 ; i < len(img.Pix) ; i += 4 {
		if i < count * 500 {
			img.Pix[i+1] = 255
			img.Pix[i+2] = 255
		}
		img.Pix[i+0] = 255
		img.Pix[i+3] = 255
	}

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

    setInterval(updateImage, 16);
}
</script>

</body>
</html>
`
