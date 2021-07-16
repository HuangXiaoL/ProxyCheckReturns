package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main()  {

	r := chi.NewMux()

	r.Get("/test/editor", requestOutput) //测试权限
}
func requestOutput(w http.ResponseWriter, r *http.Request)  {
	log.Println(w.Header())
}