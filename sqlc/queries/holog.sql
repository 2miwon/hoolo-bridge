-- name: ListHologsMostByWeek :many
SELECT place_id, COUNT(*) as mention_count
FROM public.holog
WHERE deleted_at IS NULL
  AND created_at >= NOW() - INTERVAL '7 days'
GROUP BY place_id
ORDER BY mention_count DESC
LIMIT 20;

-- name: ListHologsByPlaceId :many
SELECT id, place_id, title, content, created_at, external_url
FROM public.holog
WHERE place_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC
LIMIT $2;

-- name: GetHologByID :one
SELECT id, place_id, title, content, created_at, image_url
FROM public.holog
WHERE id = $1 AND deleted_at IS NULL;

-- name: CreateHolog :one
INSERT INTO public.holog (place_id, creator_id, schedule_id, title, content, image_url)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, place_id, creator_id, title, content, created_at;

-- name: ListHologsByUserID :many
SELECT id, place_id, creator_id, schedule_id, title, content, created_at, image_url, external_url
FROM public.holog
WHERE creator_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC;

-- TODO: TISTORY