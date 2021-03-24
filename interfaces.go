package dvras

// Gateway TODO
type Gateway interface {
    Load(DeviceID) (*Device, error)
    Save(*Device) error
}

// ApplicationService TODO
type ApplicationService interface {
    Start(command *StartCommand) error
    Stop(command *StopCommand) error
    Close() error
}
