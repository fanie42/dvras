package dvras

// StartCommand TODO
type StartCommand struct {
    DeviceID   DeviceID `json:"device_id"`
    Annotation string   `json:"annotation"`
}

// Handle TODO
// func (command *StartCommand) Handle(
//     repo Repository,
// ) error {
//     for {
//         device, err := repo.Load(command.DeviceID)
//         if err != nil {
//             return err
//         }

//         err := device.Start(command.Annotation)
//         if err != nil {
//             return err
//         }

//         err = repo.Save(device, events)
//         if err != SequenceConflictError {
//             return err
//         }
//     }
// }

// StopCommand TODO
type StopCommand struct {
    DeviceID   DeviceID `json:"device_id"`
    Annotation string   `json:"annotation"`
}

// Handle TODO
// func (command *StopCommand) Handle(
//     repo Repository,
// ) error {
//     for {
//         device, err := repo.Load(command.DeviceID)
//         if err != nil {
//             return err
//         }

//         events, err := device.Stop(command.Annotation)
//         if err != nil {
//             return err
//         }

//         err = repo.Save(device, events)
//         if err != SequenceConflictError {
//             return err
//         }
//     }
// }
