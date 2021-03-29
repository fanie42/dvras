package nats

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/fanie42/dvras"
    natsio "github.com/nats-io/nats.go"
)

type controller struct {
    nc  *natsio.Conn
    app dvras.ApplicationService
}

// New TODO
func New(
    connection *natsio.Conn,
    application dvras.ApplicationService,
) dvras.Controller {
    ctrl := &controller{
        nc:  connection,
        app: application,
    }

    return ctrl
}

// Run TODO
func (ctrl *controller) Run() {
    subStart, err := ctrl.nc.Subscribe(
        "dvras.commands.start",
        ctrl.start,
    )

    subStop, err := ctrl.nc.Subscribe(
        "dvras.commands.stop",
        ctrl.stop,
    )

    // http.HandleFunc("/api/start", ctrl.start)
    // http.HandleFunc("/api/stop", ctrl.stop)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func (ctrl *controller) start(msg *natsio.Msg) {
    dto := startCommandDTO{}
    err := json.NewDecoder(msg.Data).Decode(&dto)
    if err != nil {
        respond(err.Error(), http.StatusBadRequest, w)
        return
    }

    err = ctrl.app.Start()
    natsio.Publish(msg.Reply)
    if err != nil {
        respond(err.Error(), http.StatusConflict, w)
        return
    }

    respond("Successfully started", http.StatusOK, w)
    return
}

func (ctrl *controller) stop(msg *natsio.Msg) {
    dto := stopCommandDTO{}
    err := json.NewDecoder(msg.Data).Decode(&dto)
    if err != nil {
        respond(err.Error(), http.StatusBadRequest, w)
        return
    }

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
