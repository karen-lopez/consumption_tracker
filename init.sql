CREATE TABLE IF NOT EXISTS consumption_tracker (
    id SERIAL PRIMARY KEY,
    meter_id INT NOT NULL,
    active_energy INT NOT NULL,
    reactive_energy INT NOT NULL,
    capacitive_reactive INT NOT NULL,
    solar INT NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COPY consumption_tracker (meter_id, active_energy, reactive_energy, capacitive_reactive, solar, date)
    FROM '/docker-entrypoint-initdb.d/consumptions.csv'
    DELIMITER ','
    CSV HEADER;
