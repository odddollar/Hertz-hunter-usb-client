package utils

import (
	"fmt"
	"time"
)

// Convert array of ints to array of strings
func IntsToStrings(a []int) []string {
	out := make([]string, len(a))
	for i, v := range a {
		out[i] = fmt.Sprint(v)
	}
	return out
}

// Convert array of durations to array of formatted strings
func DurationsToStrings(a []time.Duration) []string {
	out := make([]string, len(a))
	for i, d := range a {
		out[i] = fmt.Sprintf("%.2gs", d.Seconds())
	}
	return out
}
