package timescaledb

import (
    "context"

    "github.com/fanie42/dvras/pkg/acquisition"
    "github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
    db *pgxpool.Pool
}

// New TODO
func New(
    database *pgxpool.Pool,
) acquisition.Repository {
    return &repository{
        db: database,
    }
}

// Load TODO
func (repo *repository) Load(
    id acquisition.DeviceID,
) (*acquisition.Device, error) {
    sql := `SELECT * FROM devices WHERE id=$1;`

    row := repo.db.QueryRow(context.Background(), sql, id)

    device := acquisition.Device{}
    err := row.Scan(&device)
    if err != nil {
        return nil, err
    }

    return &device, nil
}

// Save TODO
func (repo *repository) Save(
    device *acquisition.Device,
) error {
    for _, event := range device.Changes() {
        // Should be conditional
        sql := `INSERT
            INTO events (seq, timestamp, type, data)
            VALUES ($1, $2, $3, $4);`

        tag, err := repo.db.Exec(context.Background(), sql,
            event.Sequence,
            event.Timestamp,
            "dvras",
            event.Encode(),
        )
        if err != nil {
            // can't return here...
            return err
        }

        event.Apply(device)
    }

    sql := `INSERT INTO devices (id, state) 
        VALUES ($1, $2);`

    tag, err := repo.db.Exec(context.Background(), sql)
    if err != nil {
        return err
    }

    return nil
}
