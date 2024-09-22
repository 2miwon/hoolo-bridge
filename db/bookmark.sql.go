// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: bookmark.sql

package db

import (
	"context"
)

const deleteBookmarkByPlaceId = `-- name: DeleteBookmarkByPlaceId :one
DELETE FROM public.bookmark
WHERE user_id = $1 AND place_id = $2
RETURNING user_id, place_id
`

type DeleteBookmarkByPlaceIdParams struct {
	UserID  string `json:"user_id"`
	PlaceID int32  `json:"place_id"`
}

type DeleteBookmarkByPlaceIdRow struct {
	UserID  string `json:"user_id"`
	PlaceID int32  `json:"place_id"`
}

func (q *Queries) DeleteBookmarkByPlaceId(ctx context.Context, arg DeleteBookmarkByPlaceIdParams) (DeleteBookmarkByPlaceIdRow, error) {
	row := q.db.QueryRow(ctx, deleteBookmarkByPlaceId, arg.UserID, arg.PlaceID)
	var i DeleteBookmarkByPlaceIdRow
	err := row.Scan(&i.UserID, &i.PlaceID)
	return i, err
}

const getBookmarkByUserIDAndPlaceID = `-- name: GetBookmarkByUserIDAndPlaceID :one
SELECT user_id, place_id
FROM public.bookmark
WHERE user_id = $1 AND place_id = $2
`

type GetBookmarkByUserIDAndPlaceIDParams struct {
	UserID  string `json:"user_id"`
	PlaceID int32  `json:"place_id"`
}

type GetBookmarkByUserIDAndPlaceIDRow struct {
	UserID  string `json:"user_id"`
	PlaceID int32  `json:"place_id"`
}

func (q *Queries) GetBookmarkByUserIDAndPlaceID(ctx context.Context, arg GetBookmarkByUserIDAndPlaceIDParams) (GetBookmarkByUserIDAndPlaceIDRow, error) {
	row := q.db.QueryRow(ctx, getBookmarkByUserIDAndPlaceID, arg.UserID, arg.PlaceID)
	var i GetBookmarkByUserIDAndPlaceIDRow
	err := row.Scan(&i.UserID, &i.PlaceID)
	return i, err
}

const setBookmarkByPlaceId = `-- name: SetBookmarkByPlaceId :one
INSERT INTO public.bookmark (user_id, place_id)
VALUES ($1, $2)
ON CONFLICT (user_id, place_id) DO NOTHING
RETURNING user_id, place_id
`

type SetBookmarkByPlaceIdParams struct {
	UserID  string `json:"user_id"`
	PlaceID int32  `json:"place_id"`
}

type SetBookmarkByPlaceIdRow struct {
	UserID  string `json:"user_id"`
	PlaceID int32  `json:"place_id"`
}

func (q *Queries) SetBookmarkByPlaceId(ctx context.Context, arg SetBookmarkByPlaceIdParams) (SetBookmarkByPlaceIdRow, error) {
	row := q.db.QueryRow(ctx, setBookmarkByPlaceId, arg.UserID, arg.PlaceID)
	var i SetBookmarkByPlaceIdRow
	err := row.Scan(&i.UserID, &i.PlaceID)
	return i, err
}
