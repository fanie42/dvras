package inmem

// type eventstore interface {
//     Save(Aggregate) error
//     Load(ID) (Aggregate, error)
// }

// type eventbus interface {
//     Publish(Event) error
//     Subscribe(EventType) <-chan Event
// }

// type repository struct {
//     store eventstore
//     bus   eventbus
// }

// // New TODO
// func New(
//     es eventstore,
//     eb eventbus,
// ) dvras.Repository {
//     repo := &repository{
//         store: es,
//         bus:   eb,
//     }

//     return repo
// }

// // Save TODO
// func (repo *repository) Save(
//     device *dvras.Device,
// ) error {
//     err := repo.store.Save(device)
//     if err != nil {
//         return err
//     }

//     err = repo.bus.Publish(device.Events())
//     if err != nil {
//         return err
//     }

//     return nil
// }

// // Load TODO
// func (repo *repository) Load(
//     id dvras.DeviceID,
// ) (*dvras.Device, error) {
//     aggregate, err := repo.store.Load(id)
//     if err != nil {
//         return nil, err
//     }

//     repo.bus.Subscribe

//     return aggregate, nil
// }
