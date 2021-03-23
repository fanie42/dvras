package rest

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/fanie42/dvras"
)

// Controller TODO
type Controller struct {
    app dvras.Service
}

// New TODO
func New(
    application dvras.Service,
) *Controller {
    ctrl := &Controller{
        app: application,
    }

    return ctrl
}

// Run TODO
func (ctrl *Controller) Run() {
    http.Handle("/", http.FileServer(http.Dir("C:/Users/Stephanus/Desktop/sansa/web/")))
    http.HandleFunc("/start", ctrl.start)
    http.HandleFunc("/stop", ctrl.stop)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Start TODO
func (ctrl *Controller) start(
    w http.ResponseWriter,
    r *http.Request,
) {
    command := dvras.StartCommand{}
    err := json.NewDecoder(r.Body).Decode(&command)
    if err != nil {
        respond(err.Error(), http.StatusBadRequest, w)
        return
    }

    err = ctrl.app.Start(&command)
    if err != nil {
        respond(err.Error(), http.StatusConflict, w)
        return
    }

    respond("Successfully started", http.StatusOK, w)
    return
}

// Start TODO
func (ctrl *Controller) stop(
    w http.ResponseWriter,
    r *http.Request,
) {
    command := dvras.StopCommand{}
    err := json.NewDecoder(r.Body).Decode(&command)
    if err != nil {
        respond(err.Error(), http.StatusBadRequest, w)
        return
    }

    err = ctrl.app.Stop(&command)
    if err != nil {
        respond(err.Error(), http.StatusConflict, w)
        return
    }

    respond("Successfully stopped", http.StatusOK, w)
    return
}

func respond(message string, code int, w http.ResponseWriter) {
    response, _ := json.Marshal(
        map[string]string{
            "message": message,
        },
    )

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusConflict)
    w.Write(response)

    return
}
