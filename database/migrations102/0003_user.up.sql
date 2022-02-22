CREATE TABLE if not exists dishes
(
    id            SERIAL PRIMARY KEY,
    dishes_name   text                            NOT NULL,
    price         text                            NOT NULL,
    restaurant_id int references restaurants (id) not null,
    creators_id   int                             not null,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    archived_at   timestamp                DEFAULT null
);

ALTER TABLE restaurants
    ADD COLUMN archived_at timestamp DEFAULT null;

ALTER TABLE users
    ADD COLUMN archived_at timestamp DEFAULT null;