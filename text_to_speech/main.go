package main

/*

1. docker 설치

   아래 사이트에서 Windows용 설치
   https://www.docker.com/products/docker-toolbox

2. 도커용 터미널 실행

   설치한 도커 툴박스에서 "Docker Quickstart Terminal" 실행

3. docker로 text-to-speech server 실행

   도커 이미지 다운로드:
   > docker pull parente/espeakbox

   도커 실행:
   > docker run --name espeakbox -d -p 8080:8080 parente/espeakbox

   컨테이너 아이피 주소 확인:
   > docker-machine ip

4. 웹서버 실행

   > go run web_tts.go

5. 웹브라우저 실행

   http://localhost/

*/

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/tts", ttsHandler)
	http.HandleFunc("/mp3", mp3Handler)

	log.Println("listen..")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func mp3Handler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")

	resp, err := http.Get("http://192.168.99.100:8080/speech?text=" + url.QueryEscape(text))
	if err != nil {
		http.Error(w, "http get failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "audio/mp3")

	defer resp.Body.Close()

	io.Copy(w, resp.Body)
}

func ttsHandler(w http.ResponseWriter, r *http.Request) {

	text := r.URL.Query().Get("text")

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<!DOCTYPE html>
	   <html>
	   <head>
	   <title>Web TTS - text to speech</title>
	   </head>
	   <body>

	   <p>Your message: "%s"</p>
	   <audio src="/mp3?text=%s" preload controls autoplay></audio>

	   <form method="GET" action="/tts">
	   Input Message:     
	   <input type="text" id="text" name="text" placeholder="Type a message">
	   </form>
	   </body>
	   </html>`, html.EscapeString(text), url.QueryEscape(text))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, `<!DOCTYPE html>
	   <html>
	   <head>
	   <title>Web TTS - text to speech</title>
	   </head>
	   <body>
	   <form method="GET" action="/tts">
	   Input Message:     
	   <input type="text" id="text" name="text" placeholder="Type a message">
	   </form>
	   </body>
	   </html>`)
}
