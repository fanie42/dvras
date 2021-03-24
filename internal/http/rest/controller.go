package rest

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/fanie42/dvras"
)

// Controller TODO
type Controller struct {
    // OR - map[command]commandHandlers
    app dvras.ApplicationService
}

// New TODO
func New(
    application dvras.ApplicationService,
) *Controller {
    ctrl := &Controller{
        app: application,
    }

    return ctrl
}

// Run TODO
func (ctrl *Controller) Run() {
    http.Handle("/", http.FileServer(
        http.Dir("C:/Users/Stephanus/Desktop/dvras/web/")),
    )
    http.HandleFunc("/start", ctrl.start)
    http.HandleFunc("/stop", ctrl.stop)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Start TODO
func (ctrl *Controller) start(
    w http.ResponseWriter,
    r *http.Request,
) {
    dto := startDTO{}
    err := json.NewDecoder(r.Body).Decode(&dto)
    if err != nil {
        respond(err.Error(), http.StatusBadRequest, w)
        return
    }

    command := &dvras.StartCommand{
        // DeviceID:   dto.deviceID,
        Annotation: dto.annotation,
    }

    err = ctrl.app.Start(command)
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
    dto := stopDTO{}
    err := json.NewDecoder(r.Body).Decode(&dto)
    if err != nil {
        respond(err.Error(), http.StatusBadRequest, w)
        return
    }

    command := &dvras.StopCommand{
        // DeviceID:   dto.deviceID,
        Annotation: dto.annotation,
    }

    err = ctrl.app.Stop(command)
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
