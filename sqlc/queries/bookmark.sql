-- name: SetBookmarkByHologId :one
INSERT INTO public.bookmark (user_id, holog_id)
VALUES ($1, $2)
RETURNING user_id, holog_id;

-- name: DeleteBookmarkByHologId :one
DELETE FROM public.bookmark
WHERE user_id = $1 AND holog_id = $2
RETURNING user_id, holog_id;

-- name: GetBookmarkByUserIDAndPlaceID :one
SELECT b.user_id, h.place_id
FROM public.bookmark b
JOIN public.holog h ON b.holog_id = h.id
WHERE b.user_id = $1 AND h.place_id = $2;