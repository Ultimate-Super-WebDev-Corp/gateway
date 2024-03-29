create table comment (
    id serial primary key,
    product_id integer not null,
    text varchar not null,
    name varchar not null,
    source varchar not null,
    customer_id integer,
    rating integer not null,
    created_at timestamp default current_timestamp,
    foreign key (product_id) references  product(id)
);

create index ix_comment_product_id on comment (product_id);