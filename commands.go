package dvras

// StartCommand TODO
type StartCommand struct {
    DeviceID   UUID
    Annotation string `json:"annotation"`
}

// Handle TODO
func (command *StartCommand) Handle(
    repo Repository,
) {
    for {
        device, err := repo.Load(command.DeviceID)

        events := device.Start(command.Annotation)

        // All events need to be persisted atomically.
    }
}

// StopCommand TODO
type StopCommand struct {
    // DeviceID   UUID
    Annotation string `json:"annotation"`
}

// Handle TODO
func (command *StartCommand) Handle(
    repo Repository,
) {
    for {
        device, err := repo.Load(command.DeviceID)

        events := device.Start(command.Annotation)

        for i, event := range events {
            device.Apply(event)
            repo.Save(device)
        }

    }
}
