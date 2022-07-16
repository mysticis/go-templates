
-- name: CreateTask :one
INSERT INTO task (
  title, content, created_date, last_modified_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: ListTasks :many
SELECT * FROM task
ORDER BY title;

-- name: GetTask :one
SELECT * FROM task
WHERE id = $1 LIMIT 1;
