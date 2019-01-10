package main

import (
	"github.com/bigflood/gostudy/gameoflife/sim"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
)

const (
	addr   = ":8080"
	width  = 500
	height = 500
)

func main() {
	rand.Seed(1234)

	sim := sim.New(width, height)

	go sim.EncodeImages()

	http.HandleFunc("/", homeHanlder)
	http.HandleFunc("/image.png", sim.ImageHanlder)

	log.Println("listen:", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
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
