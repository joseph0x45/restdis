create table users(
  username text not null primary key,
  password text not null
);

create table sessions(
  id uuid not null primary key,
  username text not null,
  FOREIGN KEY(username) REFERENCES users(username)
);
