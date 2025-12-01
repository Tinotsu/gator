-- name: GetPostsForUser :many
SELECT posts.*
FROM posts
JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
JOIN users ON feed_follows.user_id = users.id
WHERE users.id = $1
ORDER BY posts.published_at DESC
LIMIT $2;
