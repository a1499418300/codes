package main

import "testing"

func Test_replaceAndSave(t *testing.T) {
	type args struct {
		src string
		old string
		new string
	}
	tests := []struct {
		name string
		args args
	}{
		// {"rotate_login_record", args{`D:\数据sql\rotate_login_record.sql`, "rotate_login_record", "rotate_login_record2"}},
		{"rotate_login_record", args{`C:\Users\Administrator\Desktop\rotate_login_record58.sql`, "<table_name>", "rotate_login_record"}},
		// {"rotate_login_record", args{`C:\Users\Administrator\Desktop\rotate_login_schedule.sql`, "<table_name>", "rotate_login_schedule"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replaceAndSave(tt.args.src, tt.args.old, tt.args.new)
		})
	}
}
