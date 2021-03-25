package timescaledb

import (
    "context"
    "fmt"
    "log"

    "github.com/fanie42/dvras"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v4/pgxpool"
)

type gateway struct {
    db *pgxpool.Pool
}

// New TODO
func New(
    database *pgxpool.Pool,
) dvras.Gateway {
    gw := &gateway{
        db: database,
    }

    return gw
}

// Load TODO
func (gw *gateway) Load(
    id dvras.DeviceID,
) (*dvras.Device, error) {
    sql := "INSERT INTO devices (id, state, sequence) " +
        "VALUES ($1, 'off', 0) " +
        "ON CONFLICT (id) DO NOTHING;"
    _, err := gw.db.Exec(context.Background(), sql, id)

    sql = "SELECT * FROM devices WHERE id = $1;"
    row := gw.db.QueryRow(context.Background(), sql, id)

    type result struct {
        A uuid.UUID
        B string
        C uint64
    }

    var r result

    err = row.Scan(&r.A, &r.B, &r.C)
    if err != nil {
        fmt.Printf("ERROR: %v\n", err)
        return nil, err
    }

    device := &dvras.Device{
        ID:       dvras.DeviceID(r.A),
        State:    dvras.ParseState(r.B),
        Sequence: r.C,
    }

    return device, nil
}

// Save TODO
func (gw *gateway) Save(
    device *dvras.Device,
) error {
    events := device.Changes()

    fmt.Println("starting save")

    ctx := context.Background()
    tx, err := gw.db.Begin(ctx)
    if err != nil {
        log.Fatal(err)
    }

    // Get the current sequence number of the device
    row := tx.QueryRow(ctx,
        "SELECT sequence FROM devices WHERE id = $1;",
        device.ID,
    )
    var sequence uint64
    err = row.Scan(&sequence)
    if err != nil {
        err2 := tx.Rollback(ctx)
        if err2 != nil {
            log.Fatal(err2)
        }
        return err
    }

    // Check if the sequence is up to date, return error if not
    if device.Sequence != sequence+uint64(len(events)) {
        err = tx.Rollback(ctx)
        if err != nil {
            log.Fatal(err)
        }
        return dvras.SequenceConflictError{}
    }

    // At this point we know that we can insert the new device state and events
    tag, err := tx.Exec(ctx,
        "UPDATE devices SET state = $2, sequence = $3 WHERE id = $1;",
        device.ID, device.State.String(), device.Sequence,
    )
    if err != nil || tag.RowsAffected() < 1 {
        err2 := tx.Rollback(ctx)
        if err2 != nil {
            log.Fatal(err2)
        }
        return err
    }

    // Now for the events
    for _, event := range events {
        // At this point we know that we can insert the new device state and events
        tag, err = tx.Exec(ctx,
            "INSERT INTO events (device_id, time, data) "+
                "VALUES ($1, $2, $3);",
            device.ID, event.Time(), event.JSON(),
        )
        if err != nil || tag.RowsAffected() < 1 {
            err2 := tx.Rollback(ctx)
            if err2 != nil {
                log.Fatal(err2)
            }
            return err
        }
    }

    err = tx.Commit(ctx)
    if err != nil {
        log.Fatal(err)
    }

    return nil
}
