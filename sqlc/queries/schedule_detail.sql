-- name: GetMyScheduleDetailsByScheduleId :many
SELECT id, schedule_id, place_id, title
FROM public.schedule_detail
WHERE schedule_id = $1 AND deleted_at IS NULL;

-- name: GetScheduleDetailByScheduleIdAndPlaceId :many
SELECT id, schedule_id, place_id, title
FROM public.schedule_detail
WHERE schedule_id = $1 AND place_id = $2 AND deleted_at IS NULL;

-- name: CreateScheduleDetail :one
INSERT INTO public.schedule_detail (schedule_id, place_id, title)
VALUES ($1, $2, $3)
RETURNING id, schedule_id, place_id, title;

-- name: DeleteScheduleDetail :one
UPDATE public.schedule_detail
SET deleted_at = NOW()
WHERE schedule_id = $1 AND place_id = $2
RETURNING id, schedule_id, place_id, title;