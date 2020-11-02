package common

import (
	"reflect"
	"testing"
)

func TestTrimURL(t *testing.T) {
	type args struct {
		URL string
	}
	tests := []struct {
		name string
		args args
		want []string
		API  string
	}{
		{
			name: "no slash",
			args: args{
				URL: "abc/123/4/5",
			},
			want: []string{"abc", "123", "4", "5"},
			API:  "",
		},
		{
			name: "slash",
			args: args{
				URL: "/abc/123/4/5/",
			},
			want: []string{"123", "4", "5"},
			API:  "/abc",
		},
		{
			name: "API 1",
			args: args{
				URL: "/abc/123/4/5",
			},
			want: []string{"5"},
			API:  "/abc/123/4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			APIPath = tt.API
			got := TrimURL(tt.args.URL)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrimString(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "slash inner",
			args: args{
				val: "/abc//sf//4/5////6/",
			},
			want: "/abc/sf/4/5/6",
		},
		{
			name: "pre-post no slash",
			args: args{
				val: "abc//sf//4/5/6",
			},
			want: "/abc/sf/4/5/6",
		},
		{
			name: "pre-post space",
			args: args{
				val: "   abc//sf//4/5/6   ",
			},
			want: "/abc/sf/4/5/6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimSlash(tt.args.val); got != tt.want {
				t.Errorf("TrimString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetURL(t *testing.T) {
	type args struct {
		val string
		raw bool
		who string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Empty",
			args: args{
				val: "",
				raw: false,
				who: "",
			},
			want: "/",
		},
		{
			name: "One",
			args: args{
				val: "/",
				raw: false,
				who: "",
			},
			want: "/",
		},
		{
			name: "slash",
			args: args{
				val: "///api///",
				raw: false,
				who: "",
			},
			want: "/api",
		},
		{
			name: "empty include",
			args: args{
				val: "/api/def 2",
				raw: false,
				who: "",
			},
			want: "/api/def%202",
		},
		{
			name: "different chars",
			args: args{
				val: "/api/def 2/ğş",
				raw: false,
				who: "",
			},
			want: "/api/def%202/%C4%9F%C5%9F",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetURL(tt.args.val, tt.args.raw, &tt.args.who)
			if tt.args.who != tt.want {
				t.Errorf("SetURL() = %v, want %v", tt.args.who, tt.want)
			}
		})
	}
}
