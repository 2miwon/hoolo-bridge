-- name: SetBookmarkByPlaceId :one
INSERT INTO public.bookmark (user_id, place_id)
VALUES ($1, $2)
ON CONFLICT (user_id, place_id) DO NOTHING
RETURNING user_id, place_id;

-- name: DeleteBookmarkByPlaceId :one
DELETE FROM public.bookmark
WHERE user_id = $1 AND place_id = $2
RETURNING user_id, place_id;

-- name: GetBookmarkByUserIDAndPlaceID :one
SELECT user_id, place_id
FROM public.bookmark
WHERE user_id = $1 AND place_id = $2;