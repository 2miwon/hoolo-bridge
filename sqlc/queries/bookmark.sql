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

-- name: GetBookmarkByUserIDAndHologID :one
SELECT 1
FROM schedule_detail sd
JOIN schedule s ON sd.schedule_id = s.id
JOIN bookmark b ON s.user_id = b.user_id
JOIN holog h ON sd.place_id = h.place_id AND h.id = b.holog_id
WHERE s.user_id = $2
  AND sd.place_id = $1
LIMIT 1;