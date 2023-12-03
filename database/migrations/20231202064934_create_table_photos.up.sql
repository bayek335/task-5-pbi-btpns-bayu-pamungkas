CREATE TABLE IF NOT EXISTS photos(
    id VARCHAR(40) PRIMARY KEY NOT NULL,
    title VARCHAR(64),
    caption  VARCHAR(128),
    image_url VARCHAR(128),
    user_id VARCHAR(40),
    profile_image BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,

     FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);