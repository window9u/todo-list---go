package handlers

import (
	"backend/config"
	"net/http"
)

func Register(mux *http.ServeMux, con config.Config) {
	mux.Handle("/todos/create", app{config: con, handler: createTodoHandler})
	mux.Handle("/todos/read", app{config: con, handler: readTodoHandler})
	mux.Handle("/todos/update", app{config: con, handler: updateTodoHandler})
	mux.Handle("/todos/delete", app{config: con, handler: deleteTodoHandler})
}
