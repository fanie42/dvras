package monitoring

import "github.com/fanie42/dvras"

// Projection TODO
type Projection interface {
    OnStarted(*dvras.StartedEvent) error
    OnStopped(*dvras.StoppedEvent) error
    OnDatapointAcquired(*dvras.DatapointAcquiredEvent) error
}

// EventBus TODO
type EventBus interface {
    OnStarted(func(*dvras.StartedEvent) error) error
    OnStopped(func(*dvras.StoppedEvent) error) error
    OnDatapointAcquired(func(*dvras.DatapointAcquiredEvent) error) error
}

// Repository TODO
type Repository interface {
    GetState(id DeviceID) (State, error)
}

// Service TODO - we're only going to have on func for now. We'll only access
// everything through grafana for now.
type Service interface {
    GetState(*GetStateQuery) (State, error)
}
