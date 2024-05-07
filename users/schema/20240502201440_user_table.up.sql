create table if not exists users(
  id bigint not null unique,
  firstname varchar not null,
  username varchar,
  createdat date not null
);

create table if not exists game(
  ownerId bigint not null,
  score integer DEFAULT 0,
  gasStorage integer DEFAULT 1,
  gasMining integer DEFAULT 1,
  protection integer DEFAULT 1,
  CONSTRAINT fk_game_users FOREIGN KEY(ownerId) REFERENCES users(id)
);

create table if not exists userReferral(
  inviterId bigint not null,
  referralId bigint not null,
  CONSTRAINT fk_userReferral_inviterId FOREIGN KEY(inviterId) REFERENCES users(id),
  CONSTRAINT fk_userReferral_referralId FOREIGN KEY(referralId) REFERENCES users(id)
);

create index userReferral_inviter_idx on userReferral(inviterId);