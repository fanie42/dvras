package noop

import (
    "github.com/fanie42/dvras"
)

type repository struct {
    pub dvras.Publisher
}

// NewRepository TODO
func NewRepository(
    publisher dvras.Publisher,
) dvras.Repository {
    return &repository{
        pub: publisher,
    }
}

// Save TODO
func (repo *repository) Save(data *dvras.Data) error {
    events := data.Changes()

    for _, event := range events {
        repo.pub.Publish(event)
    }

    data.Empty()
    return nil
}
