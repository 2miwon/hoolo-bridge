-- name: ListAnnounces :many
SELECT id, title, content, created_at
FROM public.announce
WHERE deleted_at IS NULL
ORDER BY created_at DESC;