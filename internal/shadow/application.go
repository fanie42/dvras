package shadow

import (
    "github.com/fanie42/dvras"
)

type application struct {
    repo dvras.Repository
}

// New TODO
func New(
    repository dvras.Repository,
) dvras.ApplicationService {
    app := &application{
        repo: repository,
    }

    return app
}

// Start TODO
func (app *application) Start() error {
    return nil
}

// Stop TODO
func (app *application) Stop() error {
    return nil
}

// Close TODO
func (app *application) Close() error {
    return nil
}
