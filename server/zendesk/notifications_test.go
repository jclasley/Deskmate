package zendesk

import (
	"reflect"
	"testing"
	"time"
)

func Test_contains(t *testing.T) {
	type args struct {
		a []string
		x string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkTag(t *testing.T) {
	type args struct {
		ticket Ticket
	}
	tests := []struct {
		name  string
		args  args
		wantN []Notify
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := checkTag(tt.args.ticket); !reflect.DeepEqual(gotN, tt.wantN) {
				t.Errorf("checkTag() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_cleanCache(t *testing.T) {
	type args struct {
		ticket Ticket
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleanCache(tt.args.ticket)
		})
	}
}

func TestUpdateCache(t *testing.T) {
	type args struct {
		ticket  Ticket
		channel string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := UpdateCache(tt.args.ticket, tt.args.channel)
			if got != tt.want {
				t.Errorf("UpdateCache() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("UpdateCache() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetNotifyType(t *testing.T) {
	type args struct {
		remain time.Duration
	}
	tests := []struct {
		name           string
		args           args
		wantNotifyType int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNotifyType := GetNotifyType(tt.args.remain); gotNotifyType != tt.wantNotifyType {
				t.Errorf("GetNotifyType() = %v, want %v", gotNotifyType, tt.wantNotifyType)
			}
		})
	}
}

func TestGetTimeRemaining(t *testing.T) {
	type args struct {
		ticket Ticket
	}
	tests := []struct {
		name       string
		args       args
		wantRemain time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRemain := GetTimeRemaining(tt.args.ticket); !reflect.DeepEqual(gotRemain, tt.wantRemain) {
				t.Errorf("GetTimeRemaining() = %v, want %v", gotRemain, tt.wantRemain)
			}
		})
	}
}

func Test_prepSLANotification(t *testing.T) {
	type args struct {
		ticket Ticket
		notify int64
		tag    string
	}
	tests := []struct {
		name             string
		args             args
		wantNotification string
		wantColor        string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNotification, gotColor := prepSLANotification(tt.args.ticket, tt.args.notify, tt.args.tag)
			if gotNotification != tt.wantNotification {
				t.Errorf("prepSLANotification() gotNotification = %v, want %v", gotNotification, tt.wantNotification)
			}
			if gotColor != tt.wantColor {
				t.Errorf("prepSLANotification() gotColor = %v, want %v", gotColor, tt.wantColor)
			}
		})
	}
}

func Test_sendUpdatedNotification(t *testing.T) {
	type args struct {
		ticket  Ticket
		channel string
		tag     string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendUpdatedNotification(tt.args.ticket, tt.args.channel, tt.args.tag)
		})
	}
}

func Test_sendNewNotification(t *testing.T) {
	type args struct {
		ticket  Ticket
		channel string
		tag     string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendNewNotification(tt.args.ticket, tt.args.channel, tt.args.tag)
		})
	}
}

func Test_sendSLANotification(t *testing.T) {
	type args struct {
		ticket  Ticket
		channel string
		tag     string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendSLANotification(tt.args.ticket, tt.args.channel, tt.args.tag)
		})
	}
}
