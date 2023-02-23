CREATE TABLE IF NOT EXISTS accessory (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    year integer NOT NULL,
    Color    text NOT NULL,
    Country  text NOT NULL,
    Price   integer NOT NULL
);
