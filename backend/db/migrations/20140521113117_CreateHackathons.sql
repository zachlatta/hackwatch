
-- +goose Up
CREATE TABLE hackathons (
  id serial not null primary key,
  created timestamp not null,
  updated timestamp not null,
  name text not null,
  website text not null,
  twitter text not null,
  facebook text not null,
  date timestamp not null,
  location text not null,
  latitude text not null,
  longitude text not null,
  approved bool not null
);


-- +goose Down
DROP TABLE hackathons;

