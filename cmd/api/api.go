package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joelewaldo/jinks-backend/service/todo"
)

type APIServer struct {
    addr string
    db *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer {
    return &APIServer{
        addr: addr,
        db: db,
    }
}

func (s *APIServer) Run() error {
    router := mux.NewRouter()
    subrouter := router.PathPrefix("/api/v1").Subrouter()
    
    todoHandler := todo.NewHandler()
    todoHandler.RegisterRoutes(subrouter)

    log.Println("Listening on", s.addr)
    
    return http.ListenAndServe(s.addr, router)
}
