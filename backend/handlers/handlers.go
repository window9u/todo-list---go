package handlers

import (
	"backend/config"
	"net/http"
)

type app struct {
	config  config.Config
	handler func(w http.ResponseWriter, r *http.Request, con config.Config)
}

func (a app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.handler(w, r, a.config)
}

func createTodoHandler(w http.ResponseWriter, r *http.Request, con config.Config) {
	
}

func readTodoHandler(w http.ResponseWriter, r *http.Request, con config.Config) {

}

func updateTodoHandler(w http.ResponseWriter, r *http.Request, con config.Config) {

}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request, con config.Config) {

}
