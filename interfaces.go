package dvras

// Repository TODO
type Repository interface {
    Load(DeviceID) (*Device, error)
    Save(*Device) error
}

// ApplicationService TODO
type ApplicationService interface {
    Start() error
    Stop() error
    Close() error
}
