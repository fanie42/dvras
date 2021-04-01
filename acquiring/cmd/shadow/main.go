package main

import (
    "log"
    "net/http"
)

func main() {
    nc, err := natsio.Connect("nats://172.18.30.100:4222")
    if err != nil {
        log.Fatalf("unable to connect to NATS server: %v\n", err)
    }

    repository := timescaledb.NewRepository()
    application := shadow.NewApplication(repository)

    // sub, err := nc.Subscribe(
    //     "dvras.acquiring.events",
    //     application.,
    // )
    // if err != nil {
    //     log.Fatalf("unable to subscribe to NATS: %v\n", err)
    // }
    // defer sub.Unsubscribe()

    api := rest.NewAPI(application)

    log.Fatal(http.ListenAndServe(":8080", api))
}
