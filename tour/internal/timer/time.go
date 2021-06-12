package timer

import "time"

var Location *time.Location

func init() {
	Location, _ = time.LoadLocation("Asia/Shanghai")
}

func GetNowTime() time.Time {
	return time.Now().In(Location)
}

func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}

func ParseInLocation(layout, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, Location)
}
