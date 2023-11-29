
CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(40) NOT NULL,
    username VARCHAR(64) NOT NULL,
    email VARCHAR(64) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    is_active BOOLEAN DEFAULT FALSE,
    activated_at TIMESTAMP,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

    PRIMARY KEY(id)
);