create table users (
  username text not null unique primary key,
  password text not null,
  rights text default 'x'
);

create table config (
  jwt_secret text not null default 'a;cfjs maklsddasd'
);

create table blacklisted_access_tokens (
  token text not null unique primary key
);

insert into users(username, password, rights)
values('root', '$2y$04$wMJp8CXmEdyCkx/IvnVV9.ZsWaWKMIB5JsYwbDhrDOShN6S98Rh0O', 'a');

insert into config DEFAULT VALUES;
