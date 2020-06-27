package resources

import "fmt"

type Scheduler struct {
	ID       int64
	Year     int64
	Month    int64
	Schedule []ScheduleItems
}

type ScheduleItems struct {
	ID         int64
	Day        int64
	StartTime  string
	FinishTime string
}

type ScheduleTime struct {
	Hour   int64
	Minute int64
}

func ScheduledTimeToString(time ScheduleTime) string {
	return fmt.Sprintf("%d:%d", time.Hour, time.Minute)
}
