package reader

import (
	"reflect"
	"testing"
)

func TestGetHandle(t *testing.T) {
	type args struct {
		val interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Map object",
			args: args{
				val: map[string]interface{}{
					"id":   100,
					"name": "xx",
				},
			},
			want:    []byte(`{"id":100,"name":"xx"}`),
			wantErr: false,
		},
		{
			name: "String",
			args: args{
				val: "abc",
			},
			want:    []byte(`"abc"`),
			wantErr: false,
		},
		{
			name: "Number",
			args: args{
				val: 123,
			},
			want:    []byte("123"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetHandle(&tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHandle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHandle() = %v, want %v", got, tt.want)
			}
		})
	}
}
