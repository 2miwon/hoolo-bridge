-- -- name: ListHologs :many
-- SELECT id, user_id, title, content, created_at
-- FROM public.holog
-- WHERE deleted_at IS NULL
-- ORDER BY RANDOM()

-- -- name: GetHologByID :one
-- SELECT id, place_id, title, content, created_at, thumbnail_url, external_url
-- FROM public.holog
-- WHERE id = $1 AND deleted_at IS NULL;

-- -- name: CreateHolog :one
