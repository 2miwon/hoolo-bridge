-- name: ListHologsMostByWeek :many
SELECT place_id, COUNT(*) as mention_count
FROM public.holog
WHERE deleted_at IS NULL
  AND created_at >= NOW() - INTERVAL '7 days'
GROUP BY place_id
ORDER BY mention_count DESC
LIMIT 20;

-- name: ListHologsByPlaceId :many
SELECT h.id, h.place_id, h.creator_id, h.schedule_id, h.title, h.content, h.created_at, h.image_url, h.external_url
FROM public.holog h
LEFT JOIN public.bookmark b ON h.id = b.holog_id AND b.user_id = $2
WHERE h.place_id = $1
  AND h.deleted_at IS NULL
  AND (b.type IS NULL OR b.type != 'hide')
ORDER BY h.created_at DESC
LIMIT $3;

-- name: GetHologByID :one
SELECT h.id, h.place_id, h.creator_id, h.schedule_id, h.title, h.content, h.created_at, h.image_url, h.external_url
FROM public.holog h
LEFT JOIN public.bookmark b ON h.id = b.holog_id AND b.user_id = $2
WHERE h.id = $1
  AND h.deleted_at IS NULL
  AND (b.type IS NULL OR b.type != 'hide');

-- name: CreateHolog :one
INSERT INTO public.holog (place_id, creator_id, schedule_id, title, content, image_url)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, place_id, creator_id, title, content, created_at;

-- name: ListHologsByUserID :many
SELECT h.id, h.place_id, h.creator_id, h.schedule_id, h.title, h.content, h.created_at, h.image_url, h.external_url
FROM public.holog h
LEFT JOIN public.bookmark b ON h.id = b.holog_id AND b.user_id = $1
WHERE h.creator_id = $1
  AND h.deleted_at IS NULL
  AND(b.type IS NULL OR b.type != 'hide')
ORDER BY h.created_at DESC;

-- name: ListHologsByUserIdPlaceId :many
SELECT h.id, h.place_id, h.creator_id, h.schedule_id, h.title, h.content, h.created_at, h.image_url, h.external_url
FROM public.holog h
LEFT JOIN public.bookmark b ON h.id = b.holog_id AND b.user_id = $1
WHERE h.creator_id = $1
  AND h.place_id = $2
  AND h.deleted_at IS NULL
  AND (b.type IS NULL OR b.type != 'hide')
ORDER BY h.created_at DESC;

-- name: ListHologsByBookmark :many
SELECT h.id, h.place_id, h.creator_id, h.schedule_id, h.title, h.content, h.created_at, h.image_url, h.external_url
FROM public.holog h
JOIN public.bookmark b ON h.id = b.holog_id
WHERE b.user_id = $1
  AND h.deleted_at IS NULL
  AND (b.type IS NULL OR b.type != 'hide')
ORDER BY h.created_at DESC;

-- name: DeleteHologByID :one
UPDATE public.holog
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, place_id, creator_id, title, content, created_at;

-- name: HideHologByID :one
INSERT INTO public.bookmark (user_id, holog_id, type)
VALUES ($1, $2, 'hide')
ON CONFLICT (user_id, holog_id)
DO UPDATE SET type = 'hide'
RETURNING id, user_id, holog_id, type;

-- TODO: TISTORY