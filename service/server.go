package main

import (
    //"./repositories"
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    //"strconv"
)

type Server struct {
    //gameRepository repositories.GameRepository
    //userGameRepository repositories.UserGameRepository
}

//func NewServer(repository *repositories.GameRepository,
//    userGameRepository *repositories.UserGameRepository) *Server {
func NewServer() *Server {
    return &Server{
        //gameRepository: *repository,
        //userGameRepository: *userGameRepository,
    }
}

func (s *Server) Run() {

    router := mux.NewRouter()
    router.HandleFunc("/api/widgets", s.getWidgets).Methods("GET")
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("../static")))

    httpServer := &http.Server{
        Addr:    ":3000",
        Handler: router,
    }

    httpServer.ListenAndServe()
}


func (s Server) getWidgets(w http.ResponseWriter, r *http.Request) {

    //startStr := r.URL.Query().Get("start")
    //if startStr == "" {
    //    startStr = "0"
    //}
    //
    //start, err := strconv.Atoi(startStr)
    //if err != nil {
    //    fmt.Fprintf(w, "Error converting start request param to integer: %s\n", err)
    //}
    //
    //endStr := r.URL.Query().Get("count")
    //if endStr == "" {
    //    endStr = "10"
    //}
    //
    //count, err := strconv.Atoi(endStr)
    //if err != nil {
    //    fmt.Fprintf(w, "Error converting count request param to integer: %s\n", err)
    //}
    //
    //filter := r.URL.Query().Get("filter")

    //if err := json.NewEncoder(w).Encode(s.userGameRepository.Get(start, count, filter)); err != nil {
    //    fmt.Fprintf(w, "Error encoding JSON response: %s", err)
    //}
    if err := json.NewEncoder(w).Encode("Hello world"); err != nil {
        fmt.Fprintf(w, "Error encoding JSON response: %s", err)
    }
}
