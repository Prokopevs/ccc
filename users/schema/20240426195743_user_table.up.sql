create table if not exists users(
  id serial not null unique,
  firstname varchar not null,
  username varchar,
  createdat date not null
);