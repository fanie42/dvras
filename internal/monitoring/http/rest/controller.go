package rest

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/fanie42/dvras/pkg/monitoring"
)

// Controller TODO
type Controller struct {
    app monitoring.Service
}

// New TODO
func New(
    application monitoring.Service,
) *Controller {
    ctrl := &Controller{
        app: application,
    }

    return ctrl
}

// Run TODO
func (ctrl *Controller) Run() {
    http.HandleFunc("/state", ctrl.getState)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

// Start TODO
func (ctrl *Controller) getState(
    w http.ResponseWriter,
    r *http.Request,
) {
    query := monitoring.GetStateQuery{}
    err := json.NewDecoder(r.Body).Decode(&query)
    // fmt.Printf("%v\n", query)
    if err != nil {
        response := map[string]string{
            "message": err.Error(),
        }
        respond(response, http.StatusBadRequest, w)
        return
    }

    state, err := ctrl.app.GetState(&query)
    if err != nil {
        response := map[string]string{
            "message": err.Error(),
        }
        respond(response, http.StatusConflict, w)
        return
    }

    response := map[string]string{
        "message": "Query successful",
        "state":   state.String(),
    }

    respond(response, http.StatusOK, w)
    return
}

func respond(response map[string]string, code int, w http.ResponseWriter) {
    r, _ := json.Marshal(response)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(r)

    return
}
