package domain

import (
	"log"
	"time"
	_ "time/tzdata"
)

var defaultTimezone *time.Location

func init() {
	var err error
	defaultTimezone, err = time.LoadLocation("Europe/Moscow")
	if err != nil {
		log.Fatalf("time.LoadLocation: %v", err)
	}
}

type Date struct {
	Year  int
	Month int
	Day   int
}

func NewDate(year int, month int, day int) Date {
	return Date{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

func FromTime(t time.Time) Date {
	year, month, day := t.In(defaultTimezone).Date()
	return Date{
		Year:  year,
		Month: int(month),
		Day:   day,
	}
}

func (d Date) ToTime() time.Time {
	return time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, defaultTimezone)
}
