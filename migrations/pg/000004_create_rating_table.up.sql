create table rating (
    id serial primary key,
    product_id integer not null,
    source varchar not null,
    rating integer not null,
    votes integer not null,
    foreign key (product_id) references  product(id),
    updated_at timestamp default current_timestamp, -- todo trigger for auto update?
    unique(product_id, source)
);

create index ix_rating_updated_at on rating (updated_at);