
create table users (

                       id SERIAL PRIMARY KEY,
                       role_id  int unique,
                       name text NOT NULL,
                       email text NOT NULL,
                       password TEXT NOT NULL ,
                       role text not null,
                       created_by text default null   ,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);



create  table addresses(
                           role_id int references users (role_id) not null,
                           address text ,
                           address_lat int NOT NULL,
                           address_lnt  int NOT NULL
);


CREATE TABLE if not exists  restaurants
(
    id         SERIAL PRIMARY KEY,
    creators_id  int references users(role_id) not null,
    restaurant_name       text NOT NULL,
    dishes_name      text NOT NULL,
    address text ,
    address_lat int NOT NULL,
    address_lng  int NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);