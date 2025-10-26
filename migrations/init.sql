CREATE SCHEMA IF NOT EXISTS simada;

SET search_path TO simada;

CREATE TABLE IF NOT EXISTS vehicle_locations
(
    vehicle_id VARCHAR(20)      NOT NULL,
    latitude   DOUBLE PRECISION NOT NULL,
    longitude  DOUBLE PRECISION NOT NULL,
    timestamp  BIGINT           NOT NULL
);

CREATE INDEX idx_vehicle_locations_vehicle_id_timestamp ON vehicle_locations (vehicle_id, timestamp);

GRANT ALL PRIVILEGES ON SCHEMA simada TO tj2025;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA simada TO tj2025;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA simada TO tj2025;
