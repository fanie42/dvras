CREATE TABLE IF NOT EXISTS device_data (
    time TIMESTAMPTZ PRIMARY KEY,
    device_id UUID,
    ch1 SMALLINT[],
    ch2 SMALLINT[],
    pps SMALLINT[],
    FOREIGN KEY (device_id) REFERENCES devices (id)
);

CREATE TABLE IF NOT EXISTS devices (
    id UUID PRIMARY KEY,
    state TEXT,
    version INTEGER
);

SELECT create_hypertable(
    'device_data',
    'time',
    if_not_exists => TRUE
);