package cache

import (
	"reflect"
	"sync"
	"testing"
)

func TestCache_GetAll(t *testing.T) {
	// m := new(map[string]interface{})
	type fields struct {
		storage sync.Map
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		key     string
		value   map[string]interface{}
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			"get获取Cache中类型为map[string]interface{}的数据",
			fields{
				*new(sync.Map),
			},
			"1",
			map[string]interface{}{
				"1": 1,
			},
			args{
				"1",
			},
			map[string]interface{}{
				"1": 1,
			},
			false,
		},
		{
			"get获取Cache中类型为nil的数据",
			fields{
				*new(sync.Map),
			},
			"1",
			nil,
			args{
				"1",
			},
			map[string]interface{}{},
			false,
		},
		{
			"get获取Cache中不存在的数据",
			fields{
				*new(sync.Map),
			},
			"2",
			map[string]interface{}{
				"1": 1,
			},
			args{
				"1",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{
				storage: tt.fields.storage,
			}
			if err := c.SetAll(tt.key, tt.value); err != nil {
				t.Logf("%v", err)
			}
			got, err := c.GetAll(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cache.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cache.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
