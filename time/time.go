package time

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Time struct {
	time.Time
}

// MarshalJSON 序列化为JSON
func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("\"\""), nil
	}
	stamp := fmt.Sprintf("\"%s\"", t.Format(time.DateTime))
	return []byte(stamp), nil
}

// UnmarshalJSON 反序列化为JSON
func (t *Time) UnmarshalJSON(data []byte) error {
	var err error
	t.Time, err = time.Parse(time.DateTime, string(data)[1:20])
	return err
}

func (t *Time) String() string {
	data, _ := json.Marshal(t)
	return string(data)
}

func (t *Time) FieldType() int {
	return 64
}

func (t *Time) SetRaw(value interface{}) error {
	switch value.(type) {
	case time.Time:
		t.Time = value.(time.Time)
	}
	return nil
}

func (t *Time) RawValue() interface{} {
	str := t.Format(time.DateTime)
	if str == "0001-01-01 00:00:00" {
		return nil
	}
	return str
}

func NowDbTime() Time {
	dbTime := Time{}
	dbTime.Time = time.Now()
	return dbTime
}
func (t *Time) AddDates(years int, months int, days int) Time {
	dbTime := Time{}
	dbTime.Time = t.Time.AddDate(years, months, days)
	return dbTime
}
func TimeParse(value string) (Time, error) {
	dbTime := Time{}
	times, err := time.Parse(time.DateTime, value)
	dbTime.Time = times
	return dbTime, err
}
func (t *Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
