package acquiring

// Service TODO - Application Service
type Service interface {
    Start(*StartCommand) error
    Stop(*StopCommand) error
}

// SessionGateway TODO - Domain Service
type SessionGateway interface {
    Start() error
    Stop() error
}
