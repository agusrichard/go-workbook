CREATE TABLE IF NOT EXISTS todos
(
    id          serial PRIMARY KEY,
    username    VARCHAR(128) NOT NULL,
    title VARCHAR(128) NOT NULL,
    description TEXT NULL DEFAULT NULL,
    deadline timestamptz NULL,
    is_important BOOLEAN NULL,
    budget_amount FLOAT NULL,
    created_at  timestamptz  NOT NULL DEFAULT Now(),
    modified_at timestamptz  NOT NULL DEFAULT Now()
)