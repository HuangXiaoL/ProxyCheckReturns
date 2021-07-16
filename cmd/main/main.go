package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strings"
)

type ReturnData struct {
	RequestAddress string `json:"request_address"`// 发起IP
	RequestHeader string `json:"request_header"`//请求头
}
func main()  {
	if err := http.ListenAndServe("0.0.0.0:8082", rout()); err != nil {
		panic("Web服务初始化.....失败")
	}

}
func rout() *chi.Mux {
	r := chi.NewMux()

	r.Get("/", requestOutput)
	return r
}

func requestOutput(w http.ResponseWriter, r *http.Request)  {
	res:=ReturnData{}
	res.RequestAddress=r.RemoteAddr

	for k,v:=range r.Header{
		res.RequestHeader=res.RequestHeader+k+"=>"+strings.Join(v, ", ")
//log.Println(k,"=>",v)
	}



}