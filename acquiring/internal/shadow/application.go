package shadow

import (
    "github.com/fanie42/dvras/internal/acquiring"
)

type application struct {
    repo acquiring.Repository
    gw   acquiring.Gateway
}

// New TODO
func New(
    repository acquiring.Repository,
    gateway acquiring.Gateway,
) acquiring.Service {
    app := &application{
        repo: repository,
        gw:   gateway,
    }

    return app
}

// Start TODO
func (app *application) Start(deviceID string) error {
    device, err := app.repo.GetDeviceByID(deviceID)
    if err != nil {
        return err
    }

    err = app.gw.Start()

    err = app.repo.SaveDevice(device)
}

// Stop TODO
func (app *application) Stop() error {
    device, err := app.repo.GetDeviceByID(deviceID)
    if err != nil {
        return err
    }

    err = app.gw.Stop()

    err = app.repo.SaveDevice(device)
}
