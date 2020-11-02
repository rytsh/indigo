package reader

import (
	"encoding/json"
	"reflect"
	"strings"
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
			want:    []byte("abc"),
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

func TestGoInner(t *testing.T) {
	type args struct {
		val     interface{}
		urlPath []string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		want1   interface{}
		want2   interface{}
		wantErr bool
	}{
		{
			name: "silly",
			args: args{
				val:     `"hello"`,
				urlPath: []string{},
			},
			want:    "hello",
			want1:   nil,
			want2:   0,
			wantErr: false,
		},
		{
			name: "One",
			args: args{
				val: `{
					"abc": "hello"
				}`,
				urlPath: []string{"abc"},
			},
			want:    "hello",
			want1:   nil,
			want2:   "abc",
			wantErr: false,
		},
		{
			name: "deep 2",
			args: args{
				val: `{
					"user": {
						"name": "Eray",
						"age": 90
					}
				}`,
				urlPath: []string{"user", "name"},
			},
			want:    "Eray",
			want1:   nil,
			want2:   "name",
			wantErr: false,
		},
		{
			name: "deep array",
			args: args{
				val: `{
					"users": [
						{
							"id": 1,
							"name": "Eray",
							"age": 90
						}
					]
				}`,
				urlPath: []string{"users", "1", "age"},
			},
			want:    float64(90),
			want1:   nil,
			want2:   "age",
			wantErr: false,
		},
		{
			name: "deep array 2",
			args: args{
				val: `{
					"users": [
						{
							"id": 1,
							"name": "Eray",
							"age": 90,
							"hobies": [
								{
									"id": 1,
									"name": "games"
								}
							]
						}
					]
				}`,
				urlPath: []string{"users", "1", "hobies", "1"},
			},
			want: map[string]interface{}{
				"id":   1,
				"name": "games",
			},
			want1:   nil,
			want2:   0, // index for delete and change operations
			wantErr: false,
		},
		{
			name: "deep array error",
			args: args{
				val: `{
					"users": [
						{
							"name": "Eray",
							"age": 90
						}
					]
				}`,
				urlPath: []string{"users", "1", "age"},
			},
			want:    nil,
			want1:   nil,
			want2:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errJSON := json.Unmarshal([]byte(strings.ReplaceAll(tt.args.val.(string), "\t", "")), &tt.args.val)
			if errJSON != nil {
				t.Errorf("GoInnter Test error = %v", errJSON)
			}
			got, got1, got2, err := GoInner(&tt.args.val, tt.args.urlPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoInner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got1 != nil && tt.want1 != nil && !reflect.DeepEqual(got, &(tt.want)) {
				t.Errorf("GoInner() got = %v, want %v", *got, tt.want)
				t.Errorf("type got = %v, type want = %v", reflect.ValueOf(*got).Kind(), reflect.ValueOf(tt.want).Kind())
			}
			if got1 != nil && tt.want1 != nil && !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GoInner() got1 = %v, want %v", *got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("GoInner() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
