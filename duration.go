package duration

import (
	"fmt"
	"time"
)

func FormatFuzzy(d time.Duration) string {
	if d > time.Hour*24*30*12 { //year
		return fmt.Sprintf("%d年", d/(time.Hour*24*30*12))
	}
	if d > time.Hour*24*30 { //Month
		return fmt.Sprintf("%d个月", d/(time.Hour*24*30))
	}
	if d > time.Hour*24 { //day
		return fmt.Sprintf("%d天", d/(time.Hour*24))
	}
	if d > time.Hour {
		return fmt.Sprintf("%d个小时", d/time.Hour)
	}
	if d > time.Minute { //minite
		return fmt.Sprintf("%d分钟", d/time.Minute)
	}
	return fmt.Sprintf("%d秒", d/time.Second)
}
