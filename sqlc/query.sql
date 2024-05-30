-- name: GetVideo :one
select *
from videos
where id = $1 limit 1;

-- name: CreateVideo :exec
insert into videos(name, data)
values ($1, $2);