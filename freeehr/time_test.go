package freeehr

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

const (
	timeTestStrDate     = "2018-03-05"
	timeTestStrDateTime = "2018-03-04T06:31:09.895+09:00"
)

var (
	timeTestLoc, _       = time.LoadLocation("Asia/Tokyo")
	timeTestTimeDate     = time.Date(2018, 3, 5, 0, 0, 0, 0, timeTestLoc)
	timeTestTimeDateTime = time.Date(2018, 3, 4, 6, 31, 9, 895000000, timeTestLoc)
)

func TestFreeeDateTime_Marshal(t *testing.T) {
	out, err := json.Marshal(FreeeDateTime{timeTestTimeDateTime})
	if err != nil {
		t.Errorf("FreeeDateTime must be marshaled: err=%v", err)
	}
	got := string(out)
	if got != fmt.Sprintf("\"%s\"", timeTestStrDateTime) {
		t.Errorf("FreeeDateTime marshaled not expected: got=%v, expected=%v", got, fmt.Sprintf("\"%s\"", timeTestStrDateTime))
	}
}

func TestFreeeDateTime_Unmarshal(t *testing.T) {
	var got FreeeDateTime
	err := json.Unmarshal([]byte(fmt.Sprintf("\"%s\"", timeTestStrDateTime)), &got)
	if err != nil {
		t.Errorf("FreeeDateTime must be unmarshaled: err=%v", err)
	}
	if !got.Equal(FreeeDateTime{timeTestTimeDateTime}) {
		t.Errorf("FreeeDateTime unmarshaled not expected: got=%v, expected=%v", got, timeTestTimeDateTime)
	}
}

func TestFreeeDate_Marshal(t *testing.T) {
	out, err := json.Marshal(FreeeDate{timeTestTimeDate})
	if err != nil {
		t.Errorf("FreeeDate must be marshaled: err=%v", err)
	}
	got := string(out)
	if got != fmt.Sprintf("\"%s\"", timeTestStrDate) {
		t.Errorf("FreeeDate marshaled not expected: got=%v, expected=%v", got, fmt.Sprintf("\"%s\"", timeTestStrDate))
	}
}

func TestFreeeDate_Unmarshal(t *testing.T) {
	var got FreeeDate
	err := json.Unmarshal([]byte(fmt.Sprintf("\"%s\"", timeTestStrDate)), &got)
	if err != nil {
		t.Errorf("FreeeDate must be unmarshaled: err=%v", err)
	}
	if !got.Equal(FreeeDate{timeTestTimeDate}) {
		t.Errorf("FreeeDate unmarshaled not expected: got=%v, expected=%v", got, timeTestTimeDate)
	}
}

func TestFreeeDate_string(t *testing.T) {
	got := fmt.Sprintf("%s", FreeeDate{timeTestTimeDate})
	if got != timeTestStrDate {
		t.Errorf("FreeeDate string not expected: got=%v, expected=%v", got, timeTestStrDate)
	}
}
