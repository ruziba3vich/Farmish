CREATE TABLE IF NOT EXISTS animals (
    id UUID PRIMARY KEY,
    name VARCHAR(255),
    weight FLOAT,
    is_hungry BOOLEAN
);

CREATE TABLE IF NOT EXISTS admins (
    id UUID PRIMARY KEY,
    admin_name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS worker (
    id UUID PRIMARY KEY,
    worker_name VARCHAR(255)
);
