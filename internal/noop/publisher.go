package noop

import "github.com/fanie42/dvras"

type publisher struct{}

// NewPublisher TODO
func NewPublisher() dvras.Publisher {
    return &publisher{}
}

// Publish TODO
func (pub *publisher) Publish(event dvras.Event) error {
    return nil
}
