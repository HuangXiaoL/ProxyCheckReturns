package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type ReturnData struct {
	RequestAddress string `json:"request_address"`// 发起IP
	RequestHeader map[string][]string `json:"request_header"`//请求头
}
func main()  {
	if err := http.ListenAndServe("0.0.0.0:8088", rout()); err != nil {
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
	res.RequestHeader=r.Header
	log.Println(res)
	re,err:=json.Marshal(res)
	if err!=nil {
		w.WriteHeader(500)
	}
	w.Write(re)



}