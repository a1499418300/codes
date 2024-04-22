package main

import (
	"testing"
)

func Test_genChatroomRobot(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test", args{"chatroom_robot.csv"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			genChatroomRobot(tt.args.path)
		})
	}
}

func Test_genChatroom(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test", args{"chatroom.csv"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			genChatroom(tt.args.path)
		})
	}
}

func Test_genRobot(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test", args{"robot.csv"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			genRobot(tt.args.path)
		})
	}
}
