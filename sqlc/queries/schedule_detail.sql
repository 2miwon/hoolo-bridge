-- -- name: GetMyScheduleDetailsByScheduleId :many
-- SELECT id, schedule_id, place_id, created_at
-- FROM public.schedule_detail
-- WHERE schedule_id = $1 AND deleted_at IS NULL

-- -- name: GetScheduleDetailByScheduleIdAndPlaceId :one
-- SELECT id, schedule_id, place_id, created_at
-- FROM public.schedule_detail
-- WHERE schedule_id = $1 AND place_id = $2 AND deleted_at IS NULL

