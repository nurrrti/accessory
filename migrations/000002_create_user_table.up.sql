CREATE TABLE IF NOT EXISTS user (
                                      id bigserial PRIMARY KEY,
                                      Username text NOT NULL,
                                      Password text NOT NULL,
);