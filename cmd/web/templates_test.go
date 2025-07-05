package main

import (
	"testing"
	"time"

	"snippetbox.xmxxmx.us/internal/assert"
)

func TestHumanDate(t *testing.T) {
	// Initialize a new time.Time value and pass it to the humanDate function.
	// tm := time.Date(2025, 7, 5, 18, 48, 0, 0, time.UTC)
	// hd := humanDate(tm)

	// Check that the output from the humanDate function is in the format we
	// expect. If it isn't what we expect, use the t.Errorf() function to
	// indicate that the test has failed and log the expected and actual
	// values.
	// if hd != "05 Jul 2025 at 18:48" {
	// 	t.Errorf("got %q; want %q", hd, "05 Jul 2025 at 18:48")
	// }

	// Create a slice of anonymous structs containing the test case name,
	// input to our humanDate() function (the tm field), and expected output
	// (the want field).
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{name: "UTC", tm: time.Date(2025, 7, 5, 18, 48, 0, 0, time.UTC), want: "05 Jul 2025 at 18:48"},
		{name: "Empty", tm: time.Time{}, want: ""},
		{name: "CET", tm: time.Date(2025, 7, 5, 18, 48, 0, 0, time.FixedZone("CET", 1*60*60)), want: "05 Jul 2025 at 17:48"},
	}

	// Loop over the test cases.
	for _, tt := range tests {
		// Use the t.Run() function to run a sub-test for each test case. The
		// first parameter to this is the name of the test (used to identify the
		// sub-test in any log output) and the second parameter is an anonymous
		// function containing the actual test for each case.
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)

			// if hd != tt.want {
			// 	t.Errorf("got %q; want %q", hd, tt.want)
			// }
			// Use the new assert.Equal() helper to compare the expected and
			// actual values.
			assert.Equal(t, hd, tt.want)
		})
	}
}
