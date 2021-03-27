package dvras

// Repository TODO
type Repository interface {
    Save(*Data) error
}

// ApplicationService TODO
type ApplicationService interface {
    Start() error
    Stop() error
    Close() error
}

// Controller TODO
type Controller interface {
    Run()
}

// Publisher TODO
type Publisher interface {
    Publish(Event) error
}
