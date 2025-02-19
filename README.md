run
```
go run cmd/app/sinarmas-test/main.go
```

dependency
```
postgres
```

schema
```
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
```
