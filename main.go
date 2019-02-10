package main

import (
	"fmt"
	"net/http"
	"context"
	"time"
	"github.com/julienschmidt/httprouter"
)

type Instance struct {
	httpServer *http.Server
	router *httprouter.Router
}

const sampleRate = 44100
const seconds = 1
const theSong = "Wow.mp3"
const addr = "http://localhost:8080/"

func CreateInstance() *Instance {
	s := &Instance{
		router: httprouter.New(),
	}

	return s
}

func (it *Instance) Start() {
	it.bindRoutes()
	it.httpServer = &http.Server{Addr: addr, Handler: it.router}
	
	err := it.httpServer.ListenAndServe()

	if err != http.ErrServerClosed {
		it.Shutdown()
	} else {
		fmt.Printf("Error: Http server stopped")
	}
}

func (it *Instance) Shutdown() {
	if it.httpServer != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
		defer cancel()

		err := it.httpServer.Shutdown(ctx)
		if err != nil {
			fmt.Printf("Shit: Failed to shutdown")
		} else {
			it.httpServer = nil
		}
	}
}

func main() {
	library := CreateLibrary("C:/Users/Raz Admin/Desktop/test")
	fmt.Printf("Library:\n%v", library.ToString())

	it := CreateInstance()
	it.Start();

	return
}