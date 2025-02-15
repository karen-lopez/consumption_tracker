CREATE TABLE IF NOT EXISTS energy_consumption (
    id UUID PRIMARY KEY,
    meter_id INT NOT NULL,
    active_energy BIGINT NOT NULL,
    reactive_energy BIGINT NOT NULL,
    capacitive_reactive BIGINT NOT NULL,
    solar BIGINT NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COPY energy_consumption (id,meter_id, active_energy, reactive_energy, capacitive_reactive, solar, date)
    FROM '/docker-entrypoint-initdb.d/consumptions.csv'
    DELIMITER ','
    CSV HEADER;
