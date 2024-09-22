-- name: GetMyScheduleDetailsByScheduleId :many
SELECT id, schedule_id, place_id
FROM public.schedule_detail
WHERE schedule_id = $1 AND deleted_at IS NULL;

-- name: GetScheduleDetailByScheduleIdAndPlaceId :many
SELECT id, schedule_id, place_id
FROM public.schedule_detail
WHERE schedule_id = $1 AND place_id = $2 AND deleted_at IS NULL;

-- name: CreateScheduleDetail :one
INSERT INTO public.schedule_detail (schedule_id, place_id)
VALUES ($1, $2)
RETURNING id, schedule_id, place_id;

-- name: DeleteScheduleDetail :one
UPDATE public.schedule_detail
SET deleted_at = NOW()
WHERE schedule_id = $1 AND place_id = $2
RETURNING id, schedule_id, place_id;