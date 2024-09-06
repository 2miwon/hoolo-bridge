-- name: CreateUser :one
INSERT INTO public.users (id, password, username, created_at)
VALUES ($1, $2, $3, CURRENT_DATE)
RETURNING id, username;

-- name: GetUserByEmailAndPassword :one
SELECT id, username, profile_image_url, created_at
FROM public.users
WHERE id = $1 AND password = $2 AND deleted_at IS NULL;

-- name: GetUserByID :one
SELECT id, username, profile_image_url, created_at
FROM public.users
WHERE id = $1 AND deleted_at IS NULL;