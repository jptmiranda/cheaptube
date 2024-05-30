create table videos
(
    id   bigserial primary key,
    name text  not null,
    data bytea not null
);