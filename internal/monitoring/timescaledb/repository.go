package timescaledb

import (
    "context"

    "github.com/fanie42/dvras/pkg/monitoring"
    "github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
    db *pgxpool.Pool
}

// NewRepository TODO
func NewRepository(
    database *pgxpool.Pool,
) monitoring.Repository {
    // file, _ := os.Open("init.sql")

    // sql := []byte{}
    // _, _ := file.Read(sql)

    // tag, err := database.Exec(ctx, sql)

    return &repository{
        db: database,
    }
}

// GetState TODO
func (repo *repository) GetState(
    id monitoring.DeviceID,
) (monitoring.State, error) {
    sql := `SELECT state
        FROM devices
        WHERE device_id = $1;`

    row := repo.db.QueryRow(context.Background(), sql, id)

    var response monitoring.State
    err := row.Scan(&response)

    return response, err
}
