// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: holog.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	null "gopkg.in/guregu/null.v4"
)

const createHolog = `-- name: CreateHolog :one
INSERT INTO public.holog (place_id, creator_id, schedule_id, title, content, image_url)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, place_id, creator_id, title, content, created_at
`

type CreateHologParams struct {
	PlaceID    string      `json:"place_id"`
	CreatorID  string      `json:"creator_id"`
	ScheduleID pgtype.UUID `json:"schedule_id"`
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	ImageUrl   *string     `json:"image_url"`
}

type CreateHologRow struct {
	ID        uuid.UUID `json:"id"`
	PlaceID   string    `json:"place_id"`
	CreatorID string    `json:"creator_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt null.Time `json:"created_at"`
}

func (q *Queries) CreateHolog(ctx context.Context, arg CreateHologParams) (CreateHologRow, error) {
	row := q.db.QueryRow(ctx, createHolog,
		arg.PlaceID,
		arg.CreatorID,
		arg.ScheduleID,
		arg.Title,
		arg.Content,
		arg.ImageUrl,
	)
	var i CreateHologRow
	err := row.Scan(
		&i.ID,
		&i.PlaceID,
		&i.CreatorID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const deleteHologByID = `-- name: DeleteHologByID :one
UPDATE public.holog
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, place_id, creator_id, title, content, created_at
`

type DeleteHologByIDRow struct {
	ID        uuid.UUID `json:"id"`
	PlaceID   string    `json:"place_id"`
	CreatorID string    `json:"creator_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt null.Time `json:"created_at"`
}

func (q *Queries) DeleteHologByID(ctx context.Context, id uuid.UUID) (DeleteHologByIDRow, error) {
	row := q.db.QueryRow(ctx, deleteHologByID, id)
	var i DeleteHologByIDRow
	err := row.Scan(
		&i.ID,
		&i.PlaceID,
		&i.CreatorID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const getHologByID = `-- name: GetHologByID :one
SELECT h.id, h.place_id, h.creator_id, h.schedule_id, h.title, h.content, h.created_at, h.image_url, h.external_url
FROM public.holog h
LEFT JOIN public.bookmark b ON h.id = b.holog_id AND b.user_id = $2
WHERE h.id = $1
  AND h.deleted_at IS NULL
  AND (b.type IS NULL OR b.type != 'hide')
`

type GetHologByIDParams struct {
	ID     uuid.UUID `json:"id"`
	UserID string    `json:"user_id"`
}

type GetHologByIDRow struct {
	ID          uuid.UUID   `json:"id"`
	PlaceID     string      `json:"place_id"`
	CreatorID   string      `json:"creator_id"`
	ScheduleID  pgtype.UUID `json:"schedule_id"`
	Title       string      `json:"title"`
	Content     string      `json:"content"`
	CreatedAt   null.Time   `json:"created_at"`
	ImageUrl    *string     `json:"image_url"`
	ExternalUrl *string     `json:"external_url"`
}

func (q *Queries) GetHologByID(ctx context.Context, arg GetHologByIDParams) (GetHologByIDRow, error) {
	row := q.db.QueryRow(ctx, getHologByID, arg.ID, arg.UserID)
	var i GetHologByIDRow
	err := row.Scan(
		&i.ID,
		&i.PlaceID,
		&i.CreatorID,
		&i.ScheduleID,
		&i.Title,
		&i.Content,
		&i.CreatedAt,
		&i.ImageUrl,
		&i.ExternalUrl,
	)
	return i, err
}

const hideHologByID = `-- name: HideHologByID :one
INSERT INTO public.bookmark (user_id, holog_id, type)
VALUES ($1, $2, 'hide')
ON CONFLICT (user_id, holog_id)
DO UPDATE SET type = 'hide'
RETURNING id, user_id, holog_id, type
`

type HideHologByIDParams struct {
	UserID  string    `json:"user_id"`
	HologID uuid.UUID `json:"holog_id"`
}

func (q *Queries) HideHologByID(ctx context.Context, arg HideHologByIDParams) (Bookmark, error) {
	row := q.db.QueryRow(ctx, hideHologByID, arg.UserID, arg.HologID)
	var i Bookmark
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.HologID,
		&i.Type,
	)
	return i, err
}

const listHologsByPlaceId = `-- name: ListHologsByPlaceId :many
SELECT h.id, h.place_id, h.creator_id, h.schedule_id, h.title, h.content, h.created_at, h.image_url, h.external_url
FROM public.holog h
LEFT JOIN public.bookmark b ON h.id = b.holog_id AND b.user_id = $2
WHERE h.place_id = $1
  AND h.deleted_at IS NULL
  AND (b.type IS NULL OR b.type != 'hide')
ORDER BY h.created_at DESC
LIMIT $3
`

type ListHologsByPlaceIdParams struct {
	PlaceID string `json:"place_id"`
	UserID  string `json:"user_id"`
	Limit   int32  `json:"limit"`
}

type ListHologsByPlaceIdRow struct {
	ID          uuid.UUID   `json:"id"`
	PlaceID     string      `json:"place_id"`
	CreatorID   string      `json:"creator_id"`
	ScheduleID  pgtype.UUID `json:"schedule_id"`
	Title       string      `json:"title"`
	Content     string      `json:"content"`
	CreatedAt   null.Time   `json:"created_at"`
	ImageUrl    *string     `json:"image_url"`
	ExternalUrl *string     `json:"external_url"`
}

func (q *Queries) ListHologsByPlaceId(ctx context.Context, arg ListHologsByPlaceIdParams) ([]ListHologsByPlaceIdRow, error) {
	rows, err := q.db.Query(ctx, listHologsByPlaceId, arg.PlaceID, arg.UserID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListHologsByPlaceIdRow{}
	for rows.Next() {
		var i ListHologsByPlaceIdRow
		if err := rows.Scan(
			&i.ID,
			&i.PlaceID,
			&i.CreatorID,
			&i.ScheduleID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.ImageUrl,
			&i.ExternalUrl,
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

const listHologsByUserID = `-- name: ListHologsByUserID :many
SELECT h.id, h.place_id, h.creator_id, h.schedule_id, h.title, h.content, h.created_at, h.image_url, h.external_url
FROM public.holog h
LEFT JOIN public.bookmark b ON h.id = b.holog_id AND b.user_id = $1
WHERE h.creator_id = $1
  AND h.deleted_at IS NULL
  AND(b.type IS NULL OR b.type != 'hide')
ORDER BY h.created_at DESC
`

type ListHologsByUserIDRow struct {
	ID          uuid.UUID   `json:"id"`
	PlaceID     string      `json:"place_id"`
	CreatorID   string      `json:"creator_id"`
	ScheduleID  pgtype.UUID `json:"schedule_id"`
	Title       string      `json:"title"`
	Content     string      `json:"content"`
	CreatedAt   null.Time   `json:"created_at"`
	ImageUrl    *string     `json:"image_url"`
	ExternalUrl *string     `json:"external_url"`
}

func (q *Queries) ListHologsByUserID(ctx context.Context, userID string) ([]ListHologsByUserIDRow, error) {
	rows, err := q.db.Query(ctx, listHologsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListHologsByUserIDRow{}
	for rows.Next() {
		var i ListHologsByUserIDRow
		if err := rows.Scan(
			&i.ID,
			&i.PlaceID,
			&i.CreatorID,
			&i.ScheduleID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.ImageUrl,
			&i.ExternalUrl,
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

const listHologsByUserIdPlaceId = `-- name: ListHologsByUserIdPlaceId :many
SELECT h.id, h.place_id, h.creator_id, h.schedule_id, h.title, h.content, h.created_at, h.image_url, h.external_url
FROM public.holog h
LEFT JOIN public.bookmark b ON h.id = b.holog_id AND b.user_id = $1
WHERE h.creator_id = $1
  AND h.place_id = $2
  AND h.deleted_at IS NULL
  AND (b.type IS NULL OR b.type != 'hide')
ORDER BY h.created_at DESC
`

type ListHologsByUserIdPlaceIdParams struct {
	UserID  string `json:"user_id"`
	PlaceID string `json:"place_id"`
}

type ListHologsByUserIdPlaceIdRow struct {
	ID          uuid.UUID   `json:"id"`
	PlaceID     string      `json:"place_id"`
	CreatorID   string      `json:"creator_id"`
	ScheduleID  pgtype.UUID `json:"schedule_id"`
	Title       string      `json:"title"`
	Content     string      `json:"content"`
	CreatedAt   null.Time   `json:"created_at"`
	ImageUrl    *string     `json:"image_url"`
	ExternalUrl *string     `json:"external_url"`
}

func (q *Queries) ListHologsByUserIdPlaceId(ctx context.Context, arg ListHologsByUserIdPlaceIdParams) ([]ListHologsByUserIdPlaceIdRow, error) {
	rows, err := q.db.Query(ctx, listHologsByUserIdPlaceId, arg.UserID, arg.PlaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListHologsByUserIdPlaceIdRow{}
	for rows.Next() {
		var i ListHologsByUserIdPlaceIdRow
		if err := rows.Scan(
			&i.ID,
			&i.PlaceID,
			&i.CreatorID,
			&i.ScheduleID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
			&i.ImageUrl,
			&i.ExternalUrl,
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

const listHologsMostByWeek = `-- name: ListHologsMostByWeek :many
SELECT place_id, COUNT(*) as mention_count
FROM public.holog
WHERE deleted_at IS NULL
  AND created_at >= NOW() - INTERVAL '7 days'
GROUP BY place_id
ORDER BY mention_count DESC
LIMIT 20
`

type ListHologsMostByWeekRow struct {
	PlaceID      string `json:"place_id"`
	MentionCount int64  `json:"mention_count"`
}

func (q *Queries) ListHologsMostByWeek(ctx context.Context) ([]ListHologsMostByWeekRow, error) {
	rows, err := q.db.Query(ctx, listHologsMostByWeek)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListHologsMostByWeekRow{}
	for rows.Next() {
		var i ListHologsMostByWeekRow
		if err := rows.Scan(&i.PlaceID, &i.MentionCount); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
