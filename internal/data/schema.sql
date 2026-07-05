-- Persistence Schema for Ghost Hive Backend

CREATE TABLE hives (
    id TEXT PRIMARY KEY,
    location_lat DOUBLE PRECISION,
    location_lon DOUBLE PRECISION,
    location_alt DOUBLE PRECISION,
    status TEXT,
    energy_level DOUBLE PRECISION,
    last_updated TIMESTAMP WITH TIME ZONE
);

CREATE TABLE interceptors (
    id TEXT PRIMARY KEY,
    hive_id TEXT REFERENCES hives(id),
    status TEXT,
    battery_level DOUBLE PRECISION,
    health DOUBLE PRECISION,
    last_updated TIMESTAMP WITH TIME ZONE
);

CREATE TABLE threats (
    id TEXT PRIMARY KEY,
    drone_type TEXT,
    location_lat DOUBLE PRECISION,
    location_lon DOUBLE PRECISION,
    location_alt DOUBLE PRECISION,
    projected_target_lat DOUBLE PRECISION,
    projected_target_lon DOUBLE PRECISION,
    speed DOUBLE PRECISION,
    threat_type TEXT,
    detected_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE missions (
    id TEXT PRIMARY KEY,
    threat_id TEXT REFERENCES threats(id),
    status TEXT,
    start_time TIMESTAMP WITH TIME ZONE,
    end_time TIMESTAMP WITH TIME ZONE,
    interception_lat DOUBLE PRECISION,
    interception_lon DOUBLE PRECISION
);

CREATE TABLE telemetry (
    id SERIAL PRIMARY KEY,
    source_id TEXT,
    location_lat DOUBLE PRECISION,
    location_lon DOUBLE PRECISION,
    battery_level DOUBLE PRECISION,
    speed DOUBLE PRECISION,
    recorded_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE audit_logs (
    id SERIAL PRIMARY KEY,
    action TEXT,
    actor TEXT,
    details JSONB,
    created_at TIMESTAMP WITH TIME ZONE
);
