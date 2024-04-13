CREATE TABLE IF NOT EXISTS items (
  id          INTEGER PRIMARY KEY,
  site        TEXT NOT NULL,
  price       TEXT NOT NULL,
  smart_time  TIMESTAMP NOT NULL,
  name        TEXT NOT NULL,
  description TEXT NOT NULL,
  nickname    TEXT NOT NULL
);
