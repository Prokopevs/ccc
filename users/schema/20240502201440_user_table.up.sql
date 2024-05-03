create table if not exists users(
  id integer not null unique,
  firstname varchar not null,
  username varchar,
  createdat date not null,
  referrals integer[]
);

create table if not exists game(
  ownerId integer not null,
  score integer DEFAULT 0,
  gasStorage integer DEFAULT 1,
  gasMining integer DEFAULT 1,
  protection integer DEFAULT 1,
  CONSTRAINT fk_game_users FOREIGN KEY(ownerId) REFERENCES users(id)
);