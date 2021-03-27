package memory

import "github.com/fanie42/dvras"

type repository struct {
    buffer []*dvras.Data
}

// New TODO
func New() dvras.Repository {
    return &repository{
        buffer: []*dvras.Data{},
    }
}

// Save TODO
func (repo *repository) Save(data *dvras.Data) error {
    repo.buffer = append(repo.buffer, data)

    return nil
}
