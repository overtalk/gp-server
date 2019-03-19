package cache

import (
	"sync"
	"testing"
)

func TestCache_SetAll(t *testing.T) {
	type fields struct {
		storage sync.Map
	}
	type args struct {
		key   string
		value map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"function test",
			fields{
				*new(sync.Map),
			},
			args{
				key: "k",
				value: map[string]interface{}{
					"1": 1,
					"2": 2,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				storage: tt.fields.storage,
			}
			if err := c.SetAll(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Cache.SetAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
