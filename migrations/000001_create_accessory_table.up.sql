CREATE TABLE IF NOT EXISTS accessory (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    year integer NOT NULL,
    color    text NOT NULL,
    country  text NOT NULL,
    crice   integer NOT NULL
);
