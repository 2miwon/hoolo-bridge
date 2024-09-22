// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: schedule.sql

package db

import (
	"context"

	"github.com/google/uuid"
	null "gopkg.in/guregu/null.v4"
)

const createSchedule = `-- name: CreateSchedule :one
INSERT INTO public.schedule (user_id, start_date, end_date)
VALUES ($1, $2, $3)
RETURNING id, user_id, start_date, end_date
`

type CreateScheduleParams struct {
	UserID    string    `json:"user_id"`
	StartDate null.Time `json:"start_date"`
	EndDate   null.Time `json:"end_date"`
}

type CreateScheduleRow struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	StartDate null.Time `json:"start_date"`
	EndDate   null.Time `json:"end_date"`
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
    AND startdate < CURRENT_DATE
    AND enddate > CURRENT_DATE
`

type GetScheduleByUserIDRow struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	StartDate null.Time `json:"start_date"`
	EndDate   null.Time `json:"end_date"`
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
