package main

import (
	"net/http"
	"net/http/pprof"
	"time"
)

const (
	pprofAddr string = ":7890"
)

func main() {
	pprofHandler := http.NewServeMux()
	pprofHandler.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	server := &http.Server{Addr: pprofAddr, Handler: pprofHandler}
	go server.ListenAndServe()
	time.Sleep(time.Hour)
}
