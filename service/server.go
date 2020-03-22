package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
)

type Widget struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Price float32 `json:"price"`
}

type Server struct {
    repository Repository
}

func NewServer() *Server {

    repository := NewRepository()

    return &Server{
        repository: *repository,
    }
}

func (s *Server) Run() {

    router := mux.NewRouter()
    router.HandleFunc("/api/widgets/{id}", s.getWidgetById).Methods("GET")
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("../static")))

    httpServer := &http.Server{
        Addr:    ":3000",
        Handler: router,
    }

    httpServer.ListenAndServe()
}


func (s Server) getWidgetById(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    id := vars["id"]

    widget := s.repository.getWidgetById(id)
    if err := json.NewEncoder(w).Encode(widget); err != nil {
        fmt.Fprintf(w, "Error encoding JSON response: %s", err)
    }
}
