package resources

import (
	"fmt"
	"strconv"
	"strings"
)

type Scheduler struct {
	ID       int64
	Year     int64
	Month    int64
	Schedule []ScheduleItems
}

type ScheduleItems struct {
	ID         int64
	Day        int64
	StartTime  ScheduleTime
	FinishTime ScheduleTime
}

type ScheduleTime struct {
	Hour   int64
	Minute int64
}

func ScheduledTimeToString(time ScheduleTime) string {
	return fmt.Sprintf("%d:%d", time.Hour, time.Minute)
}

func StringToscheduledTime(time string) ScheduleTime {
	parsedTime := strings.Split(time, ":")
	hour, _ := strconv.Atoi(parsedTime[0])
	minutes, _ := strconv.Atoi(parsedTime[1])
	return ScheduleTime{
		Hour:   int64(hour),
		Minute: int64(minutes),
	}
}
