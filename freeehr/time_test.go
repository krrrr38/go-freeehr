package freeehr

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

const (
	strDate     = "2018-03-05"
	strDateTime = "2018-03-04T06:31:09.895Z"
)

var (
	timeDate     = time.Date(2018, 3, 5, 0, 0, 0, 0, time.UTC)
	timeDateTime = time.Date(2018, 3, 4, 6, 31, 9, 895000000, time.UTC)
)

func TestFreeeDateTime_Marshal(t *testing.T) {
	out, err := json.Marshal(FreeeDateTime{timeDateTime})
	if err != nil {
		t.Errorf("FreeeDateTime must be marshaled: err=%v", err)
	}
	got := string(out)
	if got != fmt.Sprintf("\"%s\"", strDateTime) {
		t.Errorf("FreeeDateTime marshaled not expected: got=%v, expected=%v", got, fmt.Sprintf("\"%s\"", strDateTime))
	}
}

func TestFreeeDateTime_Unmarshal(t *testing.T) {
	var got FreeeDateTime
	err := json.Unmarshal([]byte(fmt.Sprintf("\"%s\"", strDateTime)), &got)
	if err != nil {
		t.Errorf("FreeeDateTime must be unmarshaled: err=%v", err)
	}
	if !got.Equal(FreeeDateTime{timeDateTime}) {
		t.Errorf("FreeeDateTime unmarshaled not expected: got=%v, expected=%v", got, timeDateTime)
	}
}

func TestFreeeDate_Marshal(t *testing.T) {
	out, err := json.Marshal(FreeeDate{timeDate})
	if err != nil {
		t.Errorf("FreeeDate must be marshaled: err=%v", err)
	}
	got := string(out)
	if got != fmt.Sprintf("\"%s\"", strDate) {
		t.Errorf("FreeeDate marshaled not expected: got=%v, expected=%v", got, fmt.Sprintf("\"%s\"", strDate))
	}
}

func TestFreeeDate_Unmarshal(t *testing.T) {
	var got FreeeDate
	err := json.Unmarshal([]byte(fmt.Sprintf("\"%s\"", strDate)), &got)
	if err != nil {
		t.Errorf("FreeeDate must be unmarshaled: err=%v", err)
	}
	if !got.Equal(FreeeDate{timeDate}) {
		t.Errorf("FreeeDate unmarshaled not expected: got=%v, expected=%v", got, timeDate)
	}
}

func TestFreeeDate_string(t *testing.T) {
	got := fmt.Sprintf("%s", FreeeDate{timeDate})
	if got != strDate {
		t.Errorf("FreeeDate string not expected: got=%v, expected=%v", got, strDate)
	}
}
