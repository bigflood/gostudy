package httptransport

import (
	"github.com/bigflood/gostudy/gogen/service"
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
    
	A int
    
	B int
    
}

type AddResp struct {
    
	Sum int
    
	Sub int
    
}

func NewAddResp(
	
	sum int,
    
	sub int,
    
) AddResp {
	return AddResp{
	
	Sum: sum,
    
	Sub: sub,
    
	}
}



type MulReq struct {
    
	A int
    
	B int
    
}

type MulResp struct {
    
	Ret int
    
}

func NewMulResp(
	
	ret int,
    
) MulResp {
	return MulResp{
	
	Ret: ret,
    
	}
}




func (ht *httpTransport) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	switch r.URL.Path {
    
	case "/Add":
		var req AddReq
		json.NewDecoder(r.Body).Decode(&req)
		
        resp := NewAddResp( ht.svc.Add(
            
            req.A,
            
            req.B,
            
            ))

		json.NewEncoder(w).Encode(&resp)
    
	case "/Mul":
		var req MulReq
		json.NewDecoder(r.Body).Decode(&req)
		
        resp := NewMulResp( ht.svc.Mul(
            
            req.A,
            
            req.B,
            
            ))

		json.NewEncoder(w).Encode(&resp)
    
	}
}
