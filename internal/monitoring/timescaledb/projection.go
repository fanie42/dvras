package timescaledb

import (
    "context"

    "github.com/fanie42/dvras"
    "github.com/fanie42/dvras/pkg/monitoring"
    "github.com/jackc/pgx/v4/pgxpool"
)

type projection struct {
    db *pgxpool.Pool
}

// NewProjection TODO
func NewProjection(
    database *pgxpool.Pool,
) monitoring.Projection {
    // file, _ := os.Open("init.sql")

    // sql := []byte{}
    // _, _ := file.Read(sql)

    // tag, err := database.Exec(ctx, sql)

    return &projection{
        db: database,
    }
}

// OnStarted TODO
func (proj *projection) OnStarted(
    event *dvras.StartedEvent,
) error {
    sql := `UPDATE devices
        SET state = 'On'
        WHERE id = $1;`

    _, err := proj.db.Exec(
        context.Background(), sql,
        // event.Version(),
        event.DeviceID,
    )

    return err
}

// OnStopped TODO
func (proj *projection) OnStopped(
    event *dvras.StoppedEvent,
) error {
    sql := `UPDATE devices
        SET state = 'Off'
        WHERE id = $1;`

    _, err := proj.db.Exec(
        context.Background(), sql,
        // event.Version(),
        event.DeviceID,
    )

    return err
}

// AddDataPoint TODO
func (proj *projection) OnDatapointAcquired(
    event *dvras.DatapointAcquiredEvent,
) error {
    sql := `INSERT INTO data (time, device_id, ch1, ch2, pps)
        VALUES ($1, $2, $3, $4, $5);`

    _, err := proj.db.Exec(
        context.Background(), sql,
        event.Timestamp,
        event.DeviceID,
        event.Channel1,
        event.Channel2,
        event.PPS,
    )
    if err != nil {
        return err
    }

    return nil
}
