package main

import (
	"testing"
	"time"
)

func TestDateDiff(t *testing.T) {
	tests := []struct {
		name          string
		input1        string
		input2        string
		expectedDays  int
		expectedHours int
	}{
		{"Solo data, differenza 1 giorno", "01/01/2024", "02/01/2024", 1, 24},
		{"Data e orario, differenza 2 ore", "01/01/2024 10:00:00", "01/01/2024 12:00:00", 0, 2},
		{"Data e orario, invertiti", "02/01/2024 12:00:00", "01/01/2024 12:00:00", 1, 24},
		{"Data vuota (oggi)", "", "01/01/2024", int(absDuration(time.Now().Truncate(24*time.Hour).Sub(parseDate("01/01/2024")))/time.Hour) / 24, int(absDuration(time.Now().Truncate(24*time.Hour).Sub(parseDate("01/01/2024"))) / time.Hour)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d1 := parseDateOrNow(tt.input1)
			d2 := parseDateOrNow(tt.input2)
			diff := d2.Sub(d1)
			if diff < 0 {
				diff = -diff
			}
			giorni := int(diff.Hours()) / 24
			ore := int(diff.Hours())
			if giorni != tt.expectedDays {
				t.Errorf("Giorni: atteso %d, ottenuto %d", tt.expectedDays, giorni)
			}
			if ore != tt.expectedHours {
				t.Errorf("Ore: atteso %d, ottenuto %d", tt.expectedHours, ore)
			}
		})
	}
}

func parseDateOrNow(s string) time.Time {
	if s == "" {
		return time.Now().Truncate(24 * time.Hour)
	}
	return parseDate(s)
}

func parseDate(s string) time.Time {
	layoutDateTime := "02/01/2006 15:04:05"
	layoutDate := "02/01/2006"
	t, err := time.Parse(layoutDateTime, s)
	if err == nil {
		return t
	}
	t, err = time.Parse(layoutDate, s)
	if err == nil {
		return t
	}
	return time.Time{}
}

func absDuration(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}
