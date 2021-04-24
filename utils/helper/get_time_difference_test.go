package helper

import (
	"testing"
	"time"
)

func TestGetIndex(t *testing.T) {
	type args struct {
		currentTime time.Time
		targetTime  time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "successfully get index",
			args: args{
				currentTime: time.Date(2022, 1, 24, 0, 0, 0, 0, time.Now().Location()),
				targetTime:  time.Date(2022, 1, 23, 0, 0, 0, 0, time.Now().Location()),
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "successfully get index",
			args: args{
				currentTime: time.Date(2022, 1, 24, 0, 0, 0, 0, time.Now().Location()),
				targetTime:  time.Date(2022, 1, 12, 0, 0, 0, 0, time.Now().Location()),
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "successfully get index",
			args: args{
				currentTime: time.Date(2022, 2, 24, 0, 0, 0, 0, time.Now().Location()),
				targetTime:  time.Date(2022, 1, 12, 0, 0, 0, 0, time.Now().Location()),
			},
			want:    7,
			wantErr: false,
		},
		{
			name: "returns error when current time is smaller than target time",
			args: args{
				currentTime: time.Date(2022, 1, 12, 0, 0, 0, 0, time.Now().Location()),
				targetTime:  time.Date(2022, 1, 24, 0, 0, 0, 0, time.Now().Location()),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetIndex(tt.args.currentTime, tt.args.targetTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
