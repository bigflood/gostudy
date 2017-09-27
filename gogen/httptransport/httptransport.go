package httptransport

import (
	"dadev/gostudy/gogen/service"
	"encoding/json"
	"net/http"
)

func New(svc service.Service) http.Handler {
	return &httpTransport{
		svc: svc,
	}
}

type httpTransport struct {
	svc service.Service
}

type AddReq struct {
	A, B int
}

type AddResp struct {
	Sum int
}

type MulReq struct {
	A, B int
}

type MulResp struct {
	Mul int
}

func (ht *httpTransport) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	switch r.URL.Path {
	case "/add":
		var req AddReq
		json.NewDecoder(r.Body).Decode(&req)
		v := ht.svc.Add(req.A, req.B)
		resp := AddResp{Sum: v}
		json.NewEncoder(w).Encode(&resp)
	case "/mul":
		var req MulReq
		json.NewDecoder(r.Body).Decode(&req)
		v := ht.svc.Mul(req.A, req.B)
		resp := MulResp{Mul: v}
		json.NewEncoder(w).Encode(&resp)
	}
}
