package acquiring

// Service TODO - Application Service
type Service interface {
    Start(*StartCommand) error
    Stop(*StopCommand) error
}

// SessionGateway TODO - Infrastructure Service
type SessionGateway interface {
    Start() error
    Stop() error
}
