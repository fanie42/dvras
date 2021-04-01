package rest

import (
    "encoding/json"
    "net/http"

    "github.com/fanie42/dvras/acquiring"
    "github.com/gorilla/mux"
)

type handler struct {
    router   *mux.Router
    acquirer acquiring.Service
}

// NewHandler TODO
func NewHandler(
    acquiringSVC acquiring.Service,
) http.Handler {
    h := &handler{
        router:   mux.NewRouter(),
        acquirer: acquiringSVC,
    }

    h.router.HandleFunc("/api/start", h.start).Methods("POST")
    h.router.HandleFunc("/api/stop", h.stop).Methods("POST")

    return h
}

func (h *handler) ServeHTTP(w http.Response, r *http.Request) {
    h.router.ServeHTTP(w, r)
}

func (h *handler) start(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

    var command acquiring.StartCommand
    err := decoder.Decode(&command)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    h.acquirer.Start()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode("Device started.")
}

func (h *handler) stop(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

    var command acquiring.StopCommand
    err := decoder.Decode(&command)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    h.acquirer.Stop()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode("Device stopped.")
}
