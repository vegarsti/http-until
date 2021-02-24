package main

import (
	"context"
	"net/http"
	"os"
	"time"
)

var interruptChannel = make(chan os.Signal, 1)

func hello(w http.ResponseWriter, req *http.Request) {
	time.Sleep(100 * time.Millisecond)
	interruptChannel <- os.Interrupt
}

func main() {
	route := "/"
	http.HandleFunc(route, hello)
	server := http.Server{Addr: ":8080"}
	go server.ListenAndServe()
	<-interruptChannel
	server.Shutdown(context.Background())
}
