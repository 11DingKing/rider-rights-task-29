package domain

import "time"

// ValidateExportRange checks an export time range for obvious misuse.
//
// A single-sided bound (only from or only to set) and the absence of both
// bounds are legitimate queries that return everything up to / from a point
// in time, so they are accepted. A zero time.Time means "unset".
//
// When both bounds are supplied, from must not be later than to: a reversed
// range (e.g. from = Aug 21, to = Aug 20) would otherwise silently yield an
// empty result set, hiding the caller's mistake. from equal to to is allowed
// as a zero-length range at a single instant.
func ValidateExportRange(from, to time.Time) error {
	if from.IsZero() || to.IsZero() {
		return nil
	}
	if from.After(to) {
		return ValidationError{
			Field:   "export_range",
			Message: "from must not be later than to",
		}
	}
	return nil
}
