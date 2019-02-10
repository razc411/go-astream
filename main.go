package main

import (
	"fmt"
	"net/http"
	"net/endpoints"
)

type Instance struct {
	httpServer *http.Server
}

const sampleRate = 44100
const seconds = 1

func CreateInstance() *Instance {
	s := &Instance{

	}

	return s
}

func (it *Instance) Start() {

	// initialize portaudio
	it.httpServer = &http.Server{Addr: addr, Handler: endpoints.Router}
	it.httpServer.ListenAndServe()

	if err != http.ErrServerClosed {

	} else {

	}
}

func (it *Instnace) Shutdown() {
	if it.httpServer != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
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
	return
}