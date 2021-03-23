package timescaledb

import (
    "context"
    "fmt"
    "strconv"

    "github.com/fanie42/dvras"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
    db *pgxpool.Pool
}

// New TODO
func New(
    database *pgxpool.Pool,
) dvras.Repository {
    return &repository{
        db: database,
    }
}

// Load TODO
func (repo *repository) Load(
    id dvras.DeviceID,
) (*dvras.Device, error) {
    sql := `SELECT * FROM devices WHERE id=$1;`

    row := repo.db.QueryRow(context.Background(), sql, id)

    device := &dvras.Device{}
    err := row.Scan(device)
    if err != nil {
        return nil, err
    }

    return device, nil
}

// Save TODO
func (repo *repository) Save(
    device *dvras.Device,
) error {
    events := device.Changes()

    sql := "BEGIN; " +
        "IF SELECT 1 FROM devices WHERE id = " + uuid.UUID(device.ID()).String() + " " +
        "AND seq = " + strconv.FormatUint(device.Sequence(), 10) + " " +
        "INSERT INTO events (id, seq, timestamp, type, data) VALUES "

    for i, event := range events {
        if i > 1 {
            sql += ", "
        }
        sql += fmt.Sprintf("(%v, %v, %v, %v, %v)",
            event.ID,
            // event.Sequence,
            event.Time,
            "dvras",
            event.Data,
        )
    }

    sql += "; " +
        "UPDATE devices " +
        "SET state = $1, sequence = $2 " +
        "WHERE id = $3; " +
        "COMMIT;"

    _, err := repo.db.Exec(context.Background(), sql,
        device.State(),
        device.Sequence(),
        device.ID(),
    )
    if err != nil {
        // return SequenceConflictError
        return err
    }

    // event.Apply(device) Device should already be updated at this point.

    return nil
}
