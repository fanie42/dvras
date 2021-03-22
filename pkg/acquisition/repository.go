package acquisition

// Repository TODO
type Repository interface {
    Load(DeviceID) (*Device, error)
    Save(*Device) error
}
