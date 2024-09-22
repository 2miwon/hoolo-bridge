// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: announce.sql

package db

import (
	"context"

	"github.com/google/uuid"
	null "gopkg.in/guregu/null.v4"
)

const listAnnounces = `-- name: ListAnnounces :many
SELECT id, title, content, created_at
FROM public.announce
WHERE deleted_at IS NULL
ORDER BY created_at DESC
`

type ListAnnouncesRow struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt null.Time `json:"created_at"`
}

func (q *Queries) ListAnnounces(ctx context.Context) ([]ListAnnouncesRow, error) {
	rows, err := q.db.Query(ctx, listAnnounces)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAnnouncesRow{}
	for rows.Next() {
		var i ListAnnouncesRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.CreatedAt,
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
