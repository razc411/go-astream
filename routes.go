package main

import (
	"fmt"
	"net/http"
)

func (s *Instance) bindRoutes() {
	s.router.HandlerFunc("POST", "/play", s.handlePlay())
	s.router.HandlerFunc("POST", "/stop", s.handleStop())
	s.router.HandlerFunc("GET", "/list", s.handleList())
}

func (s *Instance) handlePlay() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Nope")
	}
}

func (s *Instance) handleStop() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Nope")
	}
}

func (s *Instance) handleList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The Song: %s", theSong )
	}
}