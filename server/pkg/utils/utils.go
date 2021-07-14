package utils

import "time"

func GetCurrentTimeStamp() string {
	dt := time.Now()

	//Format MM-DD-YYYY::hh:mm:ss
	return dt.Format("01-01-2001::00:00:00")
}
