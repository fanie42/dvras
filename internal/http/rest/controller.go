package rest

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/fanie42/dvras"
)

type controller struct {
    app dvras.ApplicationService
}

// New TODO
func New(
    application dvras.ApplicationService,
) dvras.Controller {
    ctrl := &controller{
        app: application,
    }

    return ctrl
}

// Run TODO
func (ctrl *controller) Run() {
    // http.Handle("/", http.FileServer(
    //     http.Dir("C:/Users/Stephanus/Desktop/dvras/web/")))
    http.HandleFunc("/api/start", ctrl.start)
    http.HandleFunc("/api/stop", ctrl.stop)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Start TODO
func (ctrl *controller) start(
    w http.ResponseWriter,
    r *http.Request,
) {
    dto := startCommandDTO{}
    err := json.NewDecoder(r.Body).Decode(&dto)
    if err != nil {
        respond(err.Error(), http.StatusBadRequest, w)
        return
    }

    // command := &dvras.StartCommand{
    //     // DeviceID:   dto.deviceID,
    //     Annotation: dto.annotation,
    // }

    err = ctrl.app.Start()
    if err != nil {
        respond(err.Error(), http.StatusConflict, w)
        return
    }

    respond("Successfully started", http.StatusOK, w)
    return
}

// Start TODO
func (ctrl *controller) stop(
    w http.ResponseWriter,
    r *http.Request,
) {
    dto := stopCommandDTO{}
    err := json.NewDecoder(r.Body).Decode(&dto)
    if err != nil {
        respond(err.Error(), http.StatusBadRequest, w)
        return
    }

    // command := &dvras.StopCommand{
    //     // DeviceID:   dto.deviceID,
    //     Annotation: dto.annotation,
    // }

    err = ctrl.app.Stop()
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
