CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    passport_serie TEXT NOT NULL UNIQUE,
    passport_number TEXT NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    patronymic VARCHAR(100),
    address TEXT
);



 

CREATE TABLE IF NOT EXISTS task_progress (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(100),
    started_at TIMESTAMP,
    finished_at TIMESTAMP,
    time_difference INTERVAL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);



CREATE INDEX idx_user_id ON users(id);

CREATE INDEX idx_progress_id ON task_progress(id);


