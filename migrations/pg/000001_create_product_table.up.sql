create table product (
    id serial primary key,
    name varchar not null, -- todo unique ?
    brand varchar not null, -- todo unique ?
    description varchar not null,
    images  varchar[],
    categories varchar[],
    country varchar,
    updated_at timestamp default current_timestamp -- todo trigger for auto update?
);

create index ix_product_updated_at on product (updated_at);
