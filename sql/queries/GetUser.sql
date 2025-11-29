-- name: GetUser :one
select name from users
where name = $1;
