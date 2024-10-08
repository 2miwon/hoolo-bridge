// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: schedule.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSchedule = `-- name: CreateSchedule :one
INSERT INTO public.schedule (user_id, start_date, end_date)
VALUES ($1, $2, $3)
RETURNING id, user_id, start_date, end_date
`

type CreateScheduleParams struct {
	UserID    string    `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type CreateScheduleRow struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (q *Queries) CreateSchedule(ctx context.Context, arg CreateScheduleParams) (CreateScheduleRow, error) {
	row := q.db.QueryRow(ctx, createSchedule, arg.UserID, arg.StartDate, arg.EndDate)
	var i CreateScheduleRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const getScheduleByUserID = `-- name: GetScheduleByUserID :one
SELECT id, user_id, start_date, end_date
FROM public.schedule
WHERE user_id = $1 
    AND deleted_at IS NULL
    -- AND start_date <= CURRENT_DATE
    -- AND end_date >= CURRENT_DATE
ORDER BY created_at DESC
LIMIT 1
`

type GetScheduleByUserIDRow struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (q *Queries) GetScheduleByUserID(ctx context.Context, userID string) (GetScheduleByUserIDRow, error) {
	row := q.db.QueryRow(ctx, getScheduleByUserID, userID)
	var i GetScheduleByUserIDRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}

const updateSchedule = `-- name: UpdateSchedule :one
UPDATE public.schedule
SET start_date = $2, end_date = $3, updated_at = NOW()
WHERE id = $1
RETURNING id, user_id, start_date, end_date
`

type UpdateScheduleParams struct {
	ID        uuid.UUID `json:"id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type UpdateScheduleRow struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

func (q *Queries) UpdateSchedule(ctx context.Context, arg UpdateScheduleParams) (UpdateScheduleRow, error) {
	row := q.db.QueryRow(ctx, updateSchedule, arg.ID, arg.StartDate, arg.EndDate)
	var i UpdateScheduleRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.StartDate,
		&i.EndDate,
	)
	return i, err
}
