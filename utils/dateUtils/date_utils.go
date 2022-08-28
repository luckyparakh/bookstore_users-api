package dateUtils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	dbDateLayout  = "2006-01-02 15:04:05"
)

func GetApiNowString() string {
	return GetNow().Format(apiDateLayout)
}
func GetNow() time.Time {
	return time.Now().UTC()
}
func GetDbNowString() string {
	return GetNow().Format(dbDateLayout)
}
