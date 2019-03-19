package data

import (
	"testing"
)

func Test_getCondition(t *testing.T) {
	type args struct {
		document string
		data     Data
		where    Data
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"function test",
			args{
				"player_cup",
				map[string]interface{}{
					"id": 1,
				},
				map[string]interface{}{
					"open_id": "syj",
				},
			},
			false,
		},
		{
			"nil document",
			args{
				"",
				map[string]interface{}{
					"id": 1,
				},
				map[string]interface{}{
					"open_id": "syj",
				},
			},
			true,
		},
		{
			"len(data) < 1",
			args{
				"player_cup",
				map[string]interface{}{},
				map[string]interface{}{
					"open_id": "syj",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getCondition(tt.args.document, tt.args.data, tt.args.where); (err != nil) != tt.wantErr {
				t.Errorf("getCondition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setCondition(t *testing.T) {
	type args struct {
		document string
		column   []string
		where    Data
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"function test",
			args{
				"player_cup",
				[]string{"cup"},
				map[string]interface{}{
					"player_id": 1,
				},
			},
			false,
		},
		{
			"function test",
			args{
				"",
				[]string{"cup"},
				map[string]interface{}{
					"player_id": 1,
				},
			},
			true,
		},
		{
			"function test",
			args{
				"player_cup",
				[]string{"cup"},
				nil,
			},
			true,
		},
		{
			"function test",
			args{
				"player_cup",
				[]string{},
				map[string]interface{}{
					"player_id": 1,
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := setCondition(tt.args.document, tt.args.column, tt.args.where); (err != nil) != tt.wantErr {
				t.Errorf("setCondition() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
