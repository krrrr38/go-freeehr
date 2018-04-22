package freeehr

import (
	"time"
)

const (
	dateTimeJSONFormat = `"2006-01-02T15:04:05.000+09:00"`
	dateTimeStrFormat  = "2006-01-02T15:04:05.000+09:00"
	dateJSONFormat     = `"2006-01-02"`
	dateStrFormat      = "2006-01-02"
)

var loc, _ = time.LoadLocation("Asia/Tokyo")

// FreeeDateTime is datetime object for freee api
type FreeeDateTime struct {
	time.Time
}

// Equal represents FreeeDateTime equality
func (t FreeeDateTime) Equal(u FreeeDateTime) bool {
	return t.Time.Equal(u.Time)
}

// MarshalJSON represents FreeeDateTime json marshal process
func (t FreeeDateTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(dateTimeJSONFormat)), nil
}

// UnmarshalJSON represents FreeeDateTime json un-marshal process
func (t *FreeeDateTime) UnmarshalJSON(data []byte) (err error) {
	str := string(data)
	(*t).Time, err = time.ParseInLocation(dateTimeJSONFormat, str, loc)
	return
}

// FreeeDate is datetime object for freee api
type FreeeDate struct {
	time.Time
}

// String represents FreeeDate to string format
func (t FreeeDate) String() string {
	return t.Time.Format(dateStrFormat)
}

// Equal represents FreeeDate equality
func (t FreeeDate) Equal(u FreeeDate) bool {
	return t.Time.Equal(u.Time)
}

// MarshalJSON represents FreeeDate json marshal process
func (t FreeeDate) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(dateJSONFormat)), nil
}

// UnmarshalJSON represents FreeeDate json un-marshal process
func (t *FreeeDate) UnmarshalJSON(data []byte) (err error) {
	str := string(data)
	(*t).Time, err = time.ParseInLocation(dateJSONFormat, str, loc)
	return
}
