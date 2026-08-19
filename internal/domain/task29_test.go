package domain

import (
	"testing"
	"time"
)

func TestExportRejectsReversedRange(t *testing.T) {
	from := time.Date(2026, 8, 21, 12, 0, 0, 0, time.UTC)
	to := from.Add(-time.Hour)
	if err := ValidateExportRange(from, to); err == nil {
		t.Fatal("reversed export range was accepted")
	}
}
