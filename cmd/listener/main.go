package main

import (
    "encoding/json"
    "fmt"
    "log"
    "time"

    "github.com/fanie42/dvras"
    natsio "github.com/nats-io/nats.go"
)

func main() {
    connection, err := natsio.Connect("nats://172.18.30.100:4222")
    if err != nil {
        log.Fatalf("failed to connect to nats server: %v", err)
    }
    defer connection.Close()

    sub, err := connection.Subscribe(
        "dvras.events",
        func(msg *natsio.Msg) {
            event := &dvras.AcquiredEvent{}
            err := json.Unmarshal(msg.Data, event)
            if err != nil {
                log.Printf("failed to unmarshal event: %v", err)
            }

            fmt.Printf("%d, %d, %d\n",
                len(event.NS),
                len(event.EW),
                len(event.PPS),
            )
        },
    )
    defer sub.Unsubscribe()

    time.Sleep(time.Second * 20)
}
