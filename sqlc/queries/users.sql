-- name: CreateUser :one
INSERT INTO public.users (email, password, username, created_at)
VALUES ($1, $2, $3, CURRENT_DATE)
RETURNING id, email, password, username, profile_image_url, created_at, deleted_at;

-- name: GetUserByEmailAndPassword :one
SELECT id, email, password, username, profile_image_url, created_at, deleted_at
FROM public.users
WHERE email = $1 AND password = $2 AND deleted_at IS NULL;

-- name: GetUserByID :one
SELECT id, email, password, username, profile_image_url, created_at, deleted_at
FROM public.users
WHERE id = $1 AND deleted_at IS NULL;