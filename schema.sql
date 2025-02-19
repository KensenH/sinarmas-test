CREATE TABLE users (
  id   BIGSERIAL PRIMARY KEY,
  username text      NOT NULL,
  password text NOT NULL
);

CREATE TABLE items (
  id   			BIGSERIAL PRIMARY KEY,
  name 			text      NOT NULL,
  description 	text not null
);