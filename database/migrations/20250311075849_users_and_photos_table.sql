-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE photos (
    id SERIAL PRIMARY KEY,  -- Primary Key (Auto Increment)
    title TEXT,
    caption TEXT,
    photo_url TEXT,
    user_id INTEGER NOT NULL,  -- Foreign Key ke `users.id`
    CONSTRAINT fk_user FOREIGN KEY (user_id) 
        REFERENCES users(id) 
        ON DELETE CASCADE  -- Hapus semua foto jika user dihapus
        ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS photos;
DROP TABLE IF EXISTS users CASCADE;
-- +goose StatementEnd
