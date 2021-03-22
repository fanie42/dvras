package main

import (
    "context"
    "log"

    "github.com/fanie42/dvras/internal/monitoring"
    "github.com/fanie42/dvras/internal/monitoring/http/rest"
    "github.com/fanie42/dvras/internal/monitoring/nats"
    "github.com/fanie42/dvras/internal/monitoring/timescaledb"
    "github.com/jackc/pgx/v4/pgxpool"
    natsio "github.com/nats-io/nats.go"
)

func main() {
    dbpool, err := pgxpool.Connect(
        context.Background(),
        "postgres://postgres:admin@172.18.30.100:5432/dvras",
    )
    if err != nil {
        log.Fatal(err)
    }
    defer dbpool.Close()

    repository := timescaledb.NewRepository(dbpool)
    projection := timescaledb.NewProjection(dbpool)

    conn, err := natsio.Connect(
        "nats://172.18.30.100:4222",
    )
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    eventbus := nats.New(conn)

    app := monitoring.New(
        eventbus,
        projection,
        repository,
    )

    ctrl := rest.New(app)
    ctrl.Run()
}
