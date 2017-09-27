package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func main() {
	svr := newProxyServer()
	log.Println("listen..")
	log.Fatal(http.ListenAndServe(":8080", svr))
}

type proxyServer struct {
	mutex      sync.Mutex
	clientPool *sync.Pool
}

func newProxyServer() *proxyServer {
	var clientPool = &sync.Pool{
		New: func() interface{} {
			client := &http.Client{}
			return client
		},
	}
	return &proxyServer{
		clientPool: clientPool,
	}
}

func (svr *proxyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	client := svr.clientPool.Get().(*http.Client)
	defer svr.clientPool.Put(client)

	targetURL := new(url.URL)
	*targetURL = *req.URL
	targetURL.Scheme = "https"
	targetURL.Host = "en.wikipedia.org"

	req2, err := http.NewRequest(req.Method, targetURL.String(), req.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	copyHeader(req2.Header, req.Header)

	resp, err := client.Do(req2)
	if err != nil {
		log.Fatal(err)
		return
	}

	copyHeader(rw.Header(), resp.Header)
	contentType := resp.Header.Get("Content-Type")

	if resp.Uncompressed && strings.HasPrefix(contentType, "text/html") {
		log.Println("html..")
		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		var f func(n *html.Node)
		f = func(n *html.Node) {
			if n.Type == html.TextNode {
				n.Data = strings.Replace(n.Data, "i", "_", -1)
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)

		err = html.Render(rw, doc)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		w, err := io.Copy(rw, resp.Body)
		if err != nil {
			log.Fatal("io.Copy..", err)
			return
		}
		log.Println(w, "bytes")
	}
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
