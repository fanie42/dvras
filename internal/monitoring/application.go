package monitoring

import (
    dvras "github.com/fanie42/dvras/pkg/monitoring"
)

// Application TODO
type Application struct {
    eb   dvras.EventBus
    proj dvras.Projection
    repo dvras.Repository
}

// New TODO
func New(
    eventbus dvras.EventBus,
    projection dvras.Projection,
    repository dvras.Repository,
) *Application {
    app := &Application{
        eb:   eventbus,
        proj: projection,
        repo: repository,
    }

    eventbus.OnStarted(projection.OnStarted)
    eventbus.OnStopped(projection.OnStopped)
    eventbus.OnDatapointAcquired(projection.OnDatapointAcquired)

    return app
}

// GetState TODO
func (app *Application) GetState(
    query *dvras.GetStateQuery,
) (dvras.State, error) {
    return app.repo.GetState(query.DeviceID)
}
