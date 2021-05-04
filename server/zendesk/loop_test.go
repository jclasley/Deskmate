package zendesk

import (
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	type args struct {
		host string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Connect(tt.args.host)
		})
	}
}

func Test_iteration(t *testing.T) {
	type args struct {
		t        *time.Ticker
		interval time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iteration(tt.args.t, tt.args.interval)
		})
	}
}
