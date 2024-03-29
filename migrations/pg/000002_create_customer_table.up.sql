create table customer (
    id serial primary key,
    email varchar not null,
    password varchar not null,
    name varchar not null,
    password_id integer not null,
    unique(email)
);

select setval(pg_get_serial_sequence('customer', 'id'), 1000, false);
