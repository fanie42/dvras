package dvras

// Repository TODO
type Repository interface {
    Load(DeviceID) (*Device, error)
    Save(*Device) error
}

// Service TODO
type Service interface {
    Start(command *StartCommand) error
    Stop(command *StopCommand) error
    Close() error
}
