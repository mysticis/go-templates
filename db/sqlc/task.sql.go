// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: task.sql

package db

import (
	"context"
	"time"
)

const createTask = `-- name: CreateTask :one
INSERT INTO task (
  title, content, created_date, last_modified_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, title, content, created_date, last_modified_at, finish_date, priority, category_id, task_status_id, due_date, user_id, hide
`

type CreateTaskParams struct {
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	CreatedDate    time.Time `json:"created_date"`
	LastModifiedAt time.Time `json:"last_modified_at"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.Title,
		arg.Content,
		arg.CreatedDate,
		arg.LastModifiedAt,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.CreatedDate,
		&i.LastModifiedAt,
		&i.FinishDate,
		&i.Priority,
		&i.CategoryID,
		&i.TaskStatusID,
		&i.DueDate,
		&i.UserID,
		&i.Hide,
	)
	return i, err
}

const getTask = `-- name: GetTask :one
SELECT id, title, content, created_date, last_modified_at, finish_date, priority, category_id, task_status_id, due_date, user_id, hide FROM task
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTask(ctx context.Context, id int64) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTask, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.CreatedDate,
		&i.LastModifiedAt,
		&i.FinishDate,
		&i.Priority,
		&i.CategoryID,
		&i.TaskStatusID,
		&i.DueDate,
		&i.UserID,
		&i.Hide,
	)
	return i, err
}

const listTasks = `-- name: ListTasks :many
SELECT id, title, content, created_date, last_modified_at, finish_date, priority, category_id, task_status_id, due_date, user_id, hide FROM task
ORDER BY title
`

func (q *Queries) ListTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.CreatedDate,
			&i.LastModifiedAt,
			&i.FinishDate,
			&i.Priority,
			&i.CategoryID,
			&i.TaskStatusID,
			&i.DueDate,
			&i.UserID,
			&i.Hide,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
