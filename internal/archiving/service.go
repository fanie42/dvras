package archiving

import (
    "log"
    "os"
    "time"

    "github.com/fanie42/dvras"
)

// Data TODO
type Data struct {
    DeviceID dvras.DeviceID
    Time     time.Time
    NS       []int16
    EW       []int16
    PPS      []int16
}

// Service TODO
type Service struct {
    layout string
}

// NewService TODO
func NewService(
    path string,
    base string,
) (*Service, error) {
    service := &Service{
        layout: path.Join(path, base),
    }

    return service
}

// Archive TODO
func (service *Service) Archive(data *Data) error {
    file, err := os.OpenFile(
        data.Time.Format(service.layout),
        os.O_APPEND|os.O_CREATE|os.O_WRONLY,
        0644,
    )
    if err != nil {
        return err
    }

    _, _ := file.Write(data)

    if err := file.Close(); err != nil {
        log.Printf("failed to close file: ", err)
        return err
    }

    return nil
}
