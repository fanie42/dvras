package nats

import (
    "github.com/fanie42/dvras"
    natsio "github.com/nats-io/nats.go"
)

type publisher struct {
    conn *natsio.Conn
}

// NewPublisher TODO
func NewPublisher(
    connection *natsio.Conn,
) dvras.Publisher {
    return &publisher{
        conn: connection,
    }
}

// Publish TODO
func (pub *publisher) Publish(event dvras.Event) error {
    return pub.conn.Publish(
        "dvras.events",
        event.Data(),
    )
}
