package rest

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

// Controller TODO
type Controller struct {
    app archiving.Service
}

// New TODO
func New(
    application archiving.Service,
) *Controller {
    ctrl := &Controller{
        app: application,
    }

    return ctrl
}

// Run TODO
func (ctrl *Controller) Run() {
    http.Handle(
        "/",
        http.FileServer(
            http.Dir("C:/Users/Stephanus/Desktop/dvras/web/"),
        ),
    )
    http.HandleFunc("/archive", ctrl.getFileName)

    log.Fatal(http.ListenAndServe(":8082", nil))
}

// Start TODO
func (ctrl *Controller) getFileName(
    w http.ResponseWriter,
    r *http.Request,
) {
    query := archiving.GetFileNameQuery{}
    err := json.NewDecoder(r.Body).Decode(&query)
    fmt.Printf("%v\n", query)
    if err != nil {
        respond(err.Error(), http.StatusBadRequest, w)
        return
    }

    file, err = ctrl.app.GetFileName(&query)
    if err != nil {
        respond(err.Error(), http.StatusConflict, w)
        return
    }

    response := map[string]string{
        "message":   "Query successful",
        "file_name": file.Name.String(),
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
