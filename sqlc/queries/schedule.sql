-- name: GetScheduleByUserID :one
SELECT id, user_id, start_date, end_date
FROM public.schedule
WHERE user_id = $1 
    AND deleted_at IS NULL
    AND startdate < CURRENT_DATE
    AND enddate > CURRENT_DATE;

-- name: CreateSchedule :one
INSERT INTO public.schedule (id, user_id, start_date, end_date)
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, start_date, end_date;