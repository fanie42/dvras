package portaudio

import (
    pa "github.com/gordonklaus/portaudio"
)

type application struct {
    repo   dvras.Repository
    stream *pa.Stream
    buffer [][]int16
}

// New TODO
func New(
    repository dvras.Repository,
) dvras.ApplicationService {
    return &application{
        repo: repository,
    }
}

// Start TODO
func (app *application) Start(command *dvras.StartCommand) error {
    err := app.stream.Start()
    if err != nil {
        return err
    }
}

// Stop TODO
func (app *application) Stop(command *dvras.StopCommand) error {
    err := app.stream.Stop()
    if err != nil {
        return err
    }
}

func (app *application) callback(buffer [][]int32) {

}
