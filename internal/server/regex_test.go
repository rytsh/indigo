package server

import (
	"testing"
)

func TestSetRegexString(t *testing.T) {
	type args struct {
		value string
		who   string
	}
	type pass struct {
		ui     []string
		api    []string
		folder []string
	}
	type nopass struct {
		ui     []string
		api    []string
		folder []string
	}
	tests := []struct {
		name   string
		args   []args
		pass   pass
		nopass nopass
	}{
		{
			name: "slash",
			args: []args{
				{
					value: "/",
					who:   "UI",
				},
				{
					value: "/",
					who:   "API",
				},
				{
					value: "/",
					who:   "FOLDER",
				},
			},
			pass: pass{
				ui:     []string{"/", "/#", "/#/aaaaa", "/#/aaa/bbb/cc/d", "/#abcde"},
				api:    []string{"/", "/#", "/users", "/users/1/2", "/users?a=5"},
				folder: []string{"/", "/#", "/folder/my"},
			},
			nopass: nopass{
				ui:     []string{"/users", "/users/1"},
				api:    []string{},
				folder: []string{},
			},
		},
		{
			name: "api",
			args: []args{
				{
					value: "/api",
					who:   "UI",
				},
				{
					value: "/api",
					who:   "API",
				},
				{
					value: "/api",
					who:   "FOLDER",
				},
			},
			pass: pass{
				ui:     []string{"/api", "/api/#", "/api/#/aaaaa", "/api/#/aaa/bbb/cc/d", "/api/#abcde"},
				api:    []string{"/api", "/api/users", "/api/users/1/2", "/api/users?a=5"},
				folder: []string{"/api", "/api/users/1"},
			},
			nopass: nopass{
				ui:     []string{"/api/users", "/api/users/1", "/ap", "/apis"},
				api:    []string{"/ap", "/apis"},
				folder: []string{"/", "/ap"},
			},
		},
		{
			name: "diff api",
			args: []args{
				{
					value: "/ui",
					who:   "UI",
				},
				{
					value: "/api",
					who:   "API",
				},
				{
					value: "/",
					who:   "FOLDER",
				},
			},
			pass: pass{
				ui:     []string{"/ui", "/ui/#", "/ui/#/aaaaa", "/ui/#/aaa/bbb/cc/d", "/ui/#abcde"},
				api:    []string{"/api", "/api/users", "/api/users/1/2", "/api/users?a=5"},
				folder: []string{"/", "/ui", "/api"},
			},
			nopass: nopass{
				ui:     []string{"/api/users", "/api/users/1", "/ap", "/apis", "ui/fdsf"},
				api:    []string{"/ap", "/apis", "/ui"},
				folder: []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, arg := range tt.args {
				SetRegexString(arg.value, arg.who)
			}
			// check pass
			for _, ps := range tt.pass.ui {
				if checkAll.UI.regexCheck.MatchString(ps) == false {
					t.Errorf("SetRegexString() ui reg = [%v], url [%v]", checkAll.UI.regexCheck.String(), ps)
				}
			}
			for _, ps := range tt.pass.api {
				if checkAll.API.regexCheck.MatchString(ps) == false {
					t.Errorf("SetRegexString() api reg = [%v], url [%v]", checkAll.API.regexCheck.String(), ps)
				}
			}
			for _, ps := range tt.pass.folder {
				if checkAll.FOLDER.regexCheck.MatchString(ps) == false {
					t.Errorf("SetRegexString() folder reg = [%v], url [%v]", checkAll.FOLDER.regexCheck.String(), ps)
				}
			}
			// check nopass
			for _, nps := range tt.nopass.ui {
				if checkAll.UI.regexCheck.MatchString(nps) {
					t.Errorf("SetRegexString() ui NO reg = [%v], url [%v]", checkAll.UI.regexCheck.String(), nps)
				}
			}
			for _, nps := range tt.nopass.api {
				if checkAll.API.regexCheck.MatchString(nps) {
					t.Errorf("SetRegexString() api NO reg = [%v], url [%v]", checkAll.API.regexCheck.String(), nps)
				}
			}
			for _, nps := range tt.nopass.folder {
				if checkAll.FOLDER.regexCheck.MatchString(nps) {
					t.Errorf("SetRegexString() folder NO reg = [%v], url [%v]", checkAll.FOLDER.regexCheck.String(), nps)
				}
			}
		})
	}
}
