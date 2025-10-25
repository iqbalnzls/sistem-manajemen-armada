CREATE SCHEMA simada;

CREATE TABLE vehicle_locations
(
    vehicle_id VARCHAR(50)    NOT NULL,
    latitude   DECIMAL(10, 8) NOT NULL,
    longitude  DECIMAL(11, 8) NOT NULL,
    timestamp  BIGINT         NOT NULL
);

CREATE INDEX idx_vehicle_locations_vehicle_id_timestamp ON vehicle_locations (vehicle_id, timestamp);