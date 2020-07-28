package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "strconv"
)

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
    router.HandleFunc("/api/tasks", s.getTasks).Methods("GET")
    router.HandleFunc("/api/tasks", s.createTask).Methods("POST")
    router.HandleFunc("/api/tasks/{id}", s.getTaskById).Methods("GET")
    router.HandleFunc("/api/tasks/order/{id}", s.reorderTask).Methods("PUT")
    router.HandleFunc("/api/tasks/{id}", s.updateTask).Methods("PUT")
    router.HandleFunc("/api/tasks/{id}", s.deleteTask).Methods("DELETE")
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("../static")))

    httpServer := &http.Server{
        Addr:    ":3500",
        Handler: router,
    }

    httpServer.ListenAndServe()
}

func (s Server) createTask(w http.ResponseWriter, r *http.Request) {

    var task Task

    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    task = *s.repository.createTask(task.Label, task.Description)

    if err := json.NewEncoder(w).Encode(task); err != nil {
        fmt.Fprintf(w, "Error encoding JSON response: %s", err)
    }
}

func (s Server) deleteTask(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    id := vars["id"]

    tasks := s.repository.deleteTaskById(id)
    if tasks == nil {
        http.Error(w, "No such task", http.StatusBadRequest)
        return
    }

    if err := json.NewEncoder(w).Encode(tasks); err != nil {
        fmt.Fprintf(w, "Error encoding JSON response: %s", err)
    }
}

func (s Server) getTasks(w http.ResponseWriter, r *http.Request) {

    startStr := r.URL.Query().Get("start")
    if startStr == "" {
       startStr = "0"
    }

    offset, err := strconv.ParseUint(startStr, 10, 32)
    if err != nil {
       fmt.Fprintf(w, "Error converting start request param to integer: %s\n", err)
    }

    endStr := r.URL.Query().Get("limit")
    if endStr == "" {
       endStr = "10"
    }

    limit, err := strconv.ParseUint(endStr, 10, 32)
    if err != nil {
       fmt.Fprintf(w, "Error converting limit request param to integer: %s\n", err)
    }

    //filter := r.URL.Query().Get("filter")

    tasks := s.repository.getTasks(uint32(offset), uint32(limit))
    count := s.repository.getTaskCount()

    response := &DataPage{uint32(offset), tasks, count};

    if err := json.NewEncoder(w).Encode(response); err != nil {
        fmt.Fprintf(w, "Error encoding JSON response: %s", err)
    }
}


func (s Server) getTaskById(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    id := vars["id"]

    task := s.repository.getTaskById(id)
    if err := json.NewEncoder(w).Encode(task); err != nil {
        fmt.Fprintf(w, "Error encoding JSON response: %s", err)
    }
}

func (s Server) reorderTask(w http.ResponseWriter, r *http.Request) {

    var params TaskReorderParams

    if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    vars := mux.Vars(r)
    taskId := vars["id"]

    newTasks, err := s.repository.reorderTask(taskId, &params)
    if err != nil {
        errorMessage := fmt.Sprintf("Error reordering tasks: %v", err)
        http.Error(w, errorMessage, http.StatusBadRequest)
        return
    }

    if err := json.NewEncoder(w).Encode(newTasks); err != nil {
        fmt.Fprintf(w, "Error encoding JSON response: %s", err)
    }
}

func (s Server) updateTask(w http.ResponseWriter, r *http.Request) {

    var task Task

    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    newTask := s.repository.updateTask(&task)
    if newTask == nil {
        http.Error(w, "No such task", http.StatusBadRequest)
        return
    }

    if err := json.NewEncoder(w).Encode(newTask); err != nil {
        fmt.Fprintf(w, "Error encoding JSON response: %s", err)
    }
}
