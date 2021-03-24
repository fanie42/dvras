package dvras

// StartCommand TODO
type StartCommand struct {
    // DeviceID   DeviceID `json:"device_id"`
    Annotation string `json:"annotation"`
}

// StartCommandHandler TODO
// type StartCommandHandler func(*StartCommand) error

// StopCommand TODO
type StopCommand struct {
    // DeviceID   DeviceID `json:"device_id"`
    Annotation string `json:"annotation"`
}

// StopCommandHandler TODO
// type StopCommandHandler func(*StartCommand) error
