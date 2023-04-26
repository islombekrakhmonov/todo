create table if not exists Todo(
    id serial primary key, 
    title varchar(50) not null,
    description varchar(50) not null,
    completed boolean not null default false,
    completed_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
);