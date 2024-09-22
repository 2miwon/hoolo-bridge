-- name: GetMyScheduleDetailsByScheduleId :many
SELECT id, schedule_id, place_id, updated_at
FROM public.schedule_detail
WHERE schedule_id = $1 AND deleted_at IS NULL;

-- name: GetScheduleDetailByScheduleIdAndPlaceId :one
SELECT id, schedule_id, place_id, updated_at
FROM public.schedule_detail
WHERE schedule_id = $1 AND place_id = $2 AND deleted_at IS NULL;

-- name: CreateScheduleDetail :one
SELECT id, schedule_id, place_id, updated_at
FROM public.schedule_detail
WHERE schedule_id = $1 AND place_id = $2 AND deleted_at IS NULL;

-- name: DeleteScheduleDetail :one
UPDATE public.schedule_detail
SET deleted_at = NOW()
WHERE schedule_id = $1 AND place_id = $2
RETURNING id, schedule_id, place_id, updated_at;