package domain

import "time"

func ValidateExportRange(from, to time.Time) error {
	if from.IsZero() || to.IsZero() {
		return nil
	}
	if from.Equal(to) {
		return nil
	}
	if from.Before(to) {
		return nil
	}
	return ValidationError{Field: "to", Message: "must not be before from"}
}
