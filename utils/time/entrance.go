package time

import "time"

type Entrance struct {
}

func (Entrance) StringToTime(str string) time.Time {
	inTime, _ := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	return inTime
}
