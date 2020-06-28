package timeformat

import (
	"time"
)

const (
	LongDateType  = "2006-01-02 15:04:05"
	ShortDateType = "2006-01-02"
)

func LongDateFormat(t time.Time) string {
	return t.Format(LongDateType)
}

func ShortDateFormat(t time.Time) string {
	return t.Format(ShortDateType)
}

func ParseLongDate(s string) (t time.Time, err error) {
	return time.ParseInLocation(LongDateType, s, time.Local)
}

func ParseShortDate(s string) (t time.Time, err error) {
	return time.ParseInLocation(ShortDateType, s, time.Local)
}
