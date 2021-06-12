CREATE TABLE IF NOT EXISTS light_table
(
    id          serial PRIMARY KEY,
    field_one   VARCHAR(128) NOT NULL,
    field_two   float        NOT NULL,
    field_three VARCHAR(128) NULL,
    field_four  timestamptz  NULL
)