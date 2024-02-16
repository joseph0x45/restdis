create table users (
  id integer primary key autoincrement not null,
  username text not null unique,
  password text not null,
  can_manage_users bool not null default false,
  can_manage_tokens bool not null default false,
  is_active bool not null default true
);

create table access_keys (
);

create table logs (

);
