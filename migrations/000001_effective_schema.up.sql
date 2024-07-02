CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    passport_serie INTEGER NOT NULL UNIQUE,
    passport_number INTEGER NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    surname VARCHAR(100) NOT NULL,
    path VARCHAR(100),
    address TEXT
);

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE IF NOT EXISTS task_progress (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    task_id INTEGER NOT NULL,
    started_at TIMESTAMP NOT NULL,
    finished_at TIMESTAMP,
    time_difference INTERVAL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (task_id) REFERENCES tasks (id)
);


CREATE INDEX idx_user_id ON users(id);
CREATE INDEX idx_task_id ON tasks(id);
CREATE INDEX idx_progress_id ON task_progress(id);