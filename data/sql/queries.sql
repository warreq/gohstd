-- name: get-commandid-by-command
SELECT commandid
FROM command
WHERE commandstring=$1

-- name: insert-command
INSERT INTO command
(commandstring)
VALUES ($1)
RETURNING commandid

-- name: insert-post
INSERT INTO post
(username, title, message, "timestamp")
VALUES ($1, $2, $3, $4)
RETURNING postid

-- name: delete-post
DELETE * from post
WHERE postid=$1

-- name: get-posts
SELECT username, title, message, "timestamp"
FROM post
LIMIT 100

-- name: insert-invocation
INSERT INTO invocation
(username, commandid, exitcode, "timestamp", hostname, "user", shell, directory)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING invocationid

-- name: get-tagid-by-name
SELECT tagid
FROM tag
WHERE name=$1

-- name: insert-tag
INSERT INTO tag
(name)
VALUES ($1)
RETURNING tagid

-- name: insert-invocationtag
INSERT INTO invocationtag
(tagid, invocationid)
VALUES ($1, $2)

-- name: get-invocations-by-user
SELECT invocationid, exitcode, "timestamp", hostname,
"user", shell, directory, commandstring, tags
FROM commandhistory
WHERE username = $1
LIMIT $2

-- name: get-commands-by-user
SELECT commandstring
FROM commandhistory
WHERE username = $1
LIMIT $2

-- name: insert-user
INSERT INTO "user"
(username, email, password)
VALUES ($1, $2, $3)

-- name: get-user-by-name
SELECT email, password
FROM "user"
WHERE username = $1

-- name: update-user-image
UPDATE "user"
SET imagelocation = $1
WHERE username = $2

