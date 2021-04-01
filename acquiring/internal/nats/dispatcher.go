package nats

import (
    "log"

    natsio "github.com/nats-io/nats.go"
)

type eventbus struct {
    nc *natsio.Conn
}

// NewEventBus TODO
func NewEventBus(
    conn *natsio.Conn,
    present func([]byte),
) acquiring.EventBus {
    eb := &eventbus{
        nc: conn,
    }

    sub, err := conn.Subscribe(
        "dvras.acquiring.events",
        func(msg *natsio.Msg) {
            event := &acquiring.Event{}
            if err := unmarshal(msg.Data, event); err != nil {
                log.Printf("failed to unmarshal event: %v\n", err)
            }

            handle(event)

            return nil
        },
    )

    return eb
}
