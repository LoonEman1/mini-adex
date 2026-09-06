CREATE TABLE partners (
    id BIGSERIAL PRIMARY KEY,

    uid TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    endpoint TEXT NOT NULL,

    is_enabled BOOLEAN NOT NULL DEFAULT TRUE,

    countries TEXT[] NOT NULL DEFAULT '{}',
    device_types TEXT[] NOT NULL DEFAULT '{}',

    min_bid_floor DOUBLE PRECISION NOT NULL DEFAULT 0,
    
    blocked_categories TEXT[] NOT NULL DEFAULT '{}'
);