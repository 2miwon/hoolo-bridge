// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: schedule_detail.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createScheduleDetail = `-- name: CreateScheduleDetail :one
SELECT id, schedule_id, place_id, updated_at
FROM public.schedule_detail
WHERE schedule_id = $1 AND place_id = $2 AND deleted_at IS NULL
`

type CreateScheduleDetailParams struct {
	ScheduleID uuid.UUID `json:"schedule_id"`
	PlaceID    string    `json:"place_id"`
}

type CreateScheduleDetailRow struct {
	ID         uuid.UUID        `json:"id"`
	ScheduleID uuid.UUID        `json:"schedule_id"`
	PlaceID    string           `json:"place_id"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) CreateScheduleDetail(ctx context.Context, arg CreateScheduleDetailParams) (CreateScheduleDetailRow, error) {
	row := q.db.QueryRow(ctx, createScheduleDetail, arg.ScheduleID, arg.PlaceID)
	var i CreateScheduleDetailRow
	err := row.Scan(
		&i.ID,
		&i.ScheduleID,
		&i.PlaceID,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteScheduleDetail = `-- name: DeleteScheduleDetail :one
UPDATE public.schedule_detail
SET deleted_at = NOW()
WHERE schedule_id = $1 AND place_id = $2
RETURNING id, schedule_id, place_id, updated_at
`

type DeleteScheduleDetailParams struct {
	ScheduleID uuid.UUID `json:"schedule_id"`
	PlaceID    string    `json:"place_id"`
}

type DeleteScheduleDetailRow struct {
	ID         uuid.UUID        `json:"id"`
	ScheduleID uuid.UUID        `json:"schedule_id"`
	PlaceID    string           `json:"place_id"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) DeleteScheduleDetail(ctx context.Context, arg DeleteScheduleDetailParams) (DeleteScheduleDetailRow, error) {
	row := q.db.QueryRow(ctx, deleteScheduleDetail, arg.ScheduleID, arg.PlaceID)
	var i DeleteScheduleDetailRow
	err := row.Scan(
		&i.ID,
		&i.ScheduleID,
		&i.PlaceID,
		&i.UpdatedAt,
	)
	return i, err
}

const getMyScheduleDetailsByScheduleId = `-- name: GetMyScheduleDetailsByScheduleId :many
SELECT id, schedule_id, place_id, updated_at
FROM public.schedule_detail
WHERE schedule_id = $1 AND deleted_at IS NULL
`

type GetMyScheduleDetailsByScheduleIdRow struct {
	ID         uuid.UUID        `json:"id"`
	ScheduleID uuid.UUID        `json:"schedule_id"`
	PlaceID    string           `json:"place_id"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) GetMyScheduleDetailsByScheduleId(ctx context.Context, scheduleID uuid.UUID) ([]GetMyScheduleDetailsByScheduleIdRow, error) {
	rows, err := q.db.Query(ctx, getMyScheduleDetailsByScheduleId, scheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetMyScheduleDetailsByScheduleIdRow{}
	for rows.Next() {
		var i GetMyScheduleDetailsByScheduleIdRow
		if err := rows.Scan(
			&i.ID,
			&i.ScheduleID,
			&i.PlaceID,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getScheduleDetailByScheduleIdAndPlaceId = `-- name: GetScheduleDetailByScheduleIdAndPlaceId :one
SELECT id, schedule_id, place_id, updated_at
FROM public.schedule_detail
WHERE schedule_id = $1 AND place_id = $2 AND deleted_at IS NULL
`

type GetScheduleDetailByScheduleIdAndPlaceIdParams struct {
	ScheduleID uuid.UUID `json:"schedule_id"`
	PlaceID    string    `json:"place_id"`
}

type GetScheduleDetailByScheduleIdAndPlaceIdRow struct {
	ID         uuid.UUID        `json:"id"`
	ScheduleID uuid.UUID        `json:"schedule_id"`
	PlaceID    string           `json:"place_id"`
	UpdatedAt  pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) GetScheduleDetailByScheduleIdAndPlaceId(ctx context.Context, arg GetScheduleDetailByScheduleIdAndPlaceIdParams) (GetScheduleDetailByScheduleIdAndPlaceIdRow, error) {
	row := q.db.QueryRow(ctx, getScheduleDetailByScheduleIdAndPlaceId, arg.ScheduleID, arg.PlaceID)
	var i GetScheduleDetailByScheduleIdAndPlaceIdRow
	err := row.Scan(
		&i.ID,
		&i.ScheduleID,
		&i.PlaceID,
		&i.UpdatedAt,
	)
	return i, err
}