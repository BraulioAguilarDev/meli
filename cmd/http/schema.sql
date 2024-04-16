CREATE TABLE IF NOT EXISTS items (
  id          INTEGER PRIMARY KEY,
  site        VARCHAR(100) NOT NULL,
  price       VARCHAR(100) NOT NULL,
  smart_time  VARCHAR(100) NOT NULL,
  name        VARCHAR(100) NOT NULL,
  description VARCHAR(100) NOT NULL,
  nickname    VARCHAR(100) NOT NULL
);
