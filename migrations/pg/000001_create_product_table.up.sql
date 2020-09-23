create table product (
    id serial primary key,
    name varchar not null,
    brand varchar not null,
    description varchar not null,
    updated_at timestamp default current_timestamp -- todo might better create a trigger for auto update
);

create index ix_product_updated_at on product (updated_at);