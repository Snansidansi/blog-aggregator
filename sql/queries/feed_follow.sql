-- name: CreateFeedFollow :many
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)

SELECT 
    inserted_feed_follow.*,
    users.name AS username,
    feeds.name AS feedname
FROM inserted_feed_follow
INNER JOIN users ON users.id=inserted_feed_follow.user_id
INNER JOIN feeds ON feeds.id=inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT 
    feed_follows.*,
    users.name AS username,
    feeds.name AS feedname
FROM feed_follows
INNER JOIN users ON feed_follows.user_id=users.id
INNER JOIN feeds ON feed_follows.feed_id=feeds.id
WHERE feed_follows.user_id = $1;
