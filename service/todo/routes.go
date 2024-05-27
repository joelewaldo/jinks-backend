package todo

import (
    "net/http"
    "github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
    return &Handler{}
}

func (h* Handler) RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/todos", h.handleTodos).Methods("GET")
    router.HandleFunc("/create-todo", h.handleCreateTodo).Methods("POST")
}

func (h* Handler) handleTodos(w http.ResponseWriter, r *http.Request) {

}

func (h* Handler) handleCreateTodo(w http.ResponseWriter, r *http.Request) {

}
